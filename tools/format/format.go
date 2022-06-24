package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/andrewarchi/browser/jsonutil"
	"github.com/wspace/corpus/tools"
	"golang.org/x/sys/execabs"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s <file>...", os.Args[0])
		os.Exit(2)
	}
	hasError := false
	for _, filename := range os.Args[1:] {
		if err := formatFile(filename); err != nil {
			fmt.Fprintln(os.Stderr, err)
			hasError = true
		}
	}
	if hasError {
		os.Exit(1)
	}
}

func formatFile(filename string) error {
	unformatted, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	var jsonBuf bytes.Buffer
	in := execabs.Command("underscore", "print")
	in.Stdin = bytes.NewReader(unformatted)
	in.Stdout = &jsonBuf
	if err := in.Run(); err != nil {
		return err
	}

	var p tools.Project
	if err := jsonutil.Decode(&jsonBuf, &p); err != nil {
		return err
	}

	repoName := p.RepoName()
	submodule, submoduleCloned := "", false
	if repoName != "" {
		submodule = strings.TrimSuffix(filename, "/project.json") + "/" + repoName
		stat, err := os.Stat(submodule)
		submoduleCloned = err == nil && stat.IsDir()
	}

	if submoduleCloned && p.Date == "" {
		var dateBuf bytes.Buffer
		// Get the earliest commit date rather than the topologically-first commit
		cmd := execabs.Command("git", "-C", submodule, "log", "--reverse", "--format=%ai|%ci")
		cmd.Stdout = &dateBuf
		if err := cmd.Run(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		} else {
			log := dateBuf.String()
			if i := strings.IndexByte(log, '\n'); i != -1 {
				log := log[:i]
				if j := strings.IndexByte(log, '|'); j != -1 {
					authorDate := log[:j]
					committerDate := log[j+1:]
					if authorDate == committerDate {
						p.Date = authorDate
					}
				}
			}
		}
	}

	if submoduleCloned && p.License == "" {
		if l, err := p.GetLicense(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		} else if l != "" {
			p.License = l
		} else {
			p.License = "not found"
		}
	}

	jsonBuf.Reset()
	e := json.NewEncoder(&jsonBuf)
	e.SetEscapeHTML(false)
	if err := e.Encode(p); err != nil {
		return err
	}

	var formattedBuf bytes.Buffer
	out := execabs.Command("underscore", "print")
	out.Stdin = &jsonBuf
	out.Stdout = &formattedBuf
	if err := out.Run(); err != nil {
		return err
	}

	if !bytes.Equal(unformatted, formattedBuf.Bytes()) {
		fmt.Printf("Formatted %s\n", filename)
		return os.WriteFile(filename, formattedBuf.Bytes(), 0o644)
	}
	return nil
}
