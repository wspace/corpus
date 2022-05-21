package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strings"

	"github.com/wspace/corpus/tools"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s <file>...", os.Args[0])
		os.Exit(2)
	}
	projects, err := tools.ReadProjects(os.Args[1:])
	try(err)
	for _, p := range projects {
		if p.Build != nil {
			slash := strings.IndexByte(p.ID, '/')
			shortID := p.ID[slash+1:]
			if strings.IndexByte(shortID, '-') == -1 {
				shortID += "-" + p.ID[:slash]
			}
			if len(p.Source) != 0 && strings.HasPrefix(p.Source[0], "https://github.com/wspace/") {
				if id := strings.TrimPrefix(p.Source[0], "https://github.com/wspace/"); id != shortID {
					fmt.Fprintf(os.Stdout, "%s: short ID %s does not match %s\n", p.ID, shortID, p.Source[0])
				}
			}
			b := new(bytes.Buffer)
			fmt.Fprintf(b, "FROM %s\n\n", p.Build.BaseImage)
			fmt.Fprintf(b, "RUN git clone %s /%s\n", p.Source[0], shortID)
			fmt.Fprintf(b, "WORKDIR /%s\n", shortID)
			for _, cmd := range p.Build.Setup {
				fmt.Fprintf(b, "RUN %s\n", cmd)
			}
			for _, target := range p.Build.Targets {
				fmt.Fprintf(b, "RUN %s\n", target.Build)
			}
			try(writeIfChanged(b, p.ID+".dockerfile"))
		}
	}
}

func writeIfChanged(b *bytes.Buffer, filename string) error {
	contents, err := os.ReadFile(filename)
	if (err == nil && !bytes.Equal(contents, b.Bytes())) || errors.Is(err, fs.ErrNotExist) {
		fmt.Printf("Generated %s\n", filename)
		return os.WriteFile(filename, b.Bytes(), 0o644)
	}
	return err
}

func try(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
