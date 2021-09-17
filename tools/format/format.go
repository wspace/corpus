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
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s <file>", os.Args[0])
		os.Exit(2)
	}
	filename := os.Args[1]

	var jsonBuf bytes.Buffer
	in := execabs.Command("underscore", "print", "-i", filename)
	in.Stdout = &jsonBuf
	try(in.Run())

	var p tools.Project
	try(jsonutil.Decode(&jsonBuf, &p))

	if p.Date == "" {
		submodule := strings.TrimSuffix(filename, ".json")
		if stat, err := os.Stat(submodule); err == nil && stat.IsDir() {
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
	}

	if p.License == "" {
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
	try(e.Encode(p))

	f, err := os.Create(filename + "~")
	try(err)
	out := execabs.Command("underscore", "print")
	out.Stdin = &jsonBuf
	out.Stdout = f
	try(out.Run())
	try(os.Rename(filename+"~", filename))
}

func try(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
