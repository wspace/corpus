package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"
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
	var errs []error
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
			inWorkdir, workdir := false, ""
			if len(p.Build.Setup) != 0 {
				fmt.Fprintf(b, "WORKDIR /%s\n", shortID)
				inWorkdir = true
			}
			for _, cmd := range p.Build.Setup {
				fmt.Fprintf(b, "RUN %s\n", cmd)
			}
			for _, target := range p.Build.Targets {
				if !inWorkdir || workdir != target.Workdir {
					fmt.Fprintf(b, "WORKDIR %s\n", path.Join("/", shortID, target.Workdir))
					inWorkdir, workdir = true, target.Workdir
				}
				fmt.Fprintf(b, "RUN %s\n", target.Build)
				b.WriteString("# builds: ")
				for i, binary := range target.Binaries {
					if i != 0 {
						b.WriteString(", ")
					}
					b.WriteString(path.Join("/", shortID, binary))
				}
				b.WriteByte('\n')
			}
			try(writeIfChanged(b, p.ID+".dockerfile"))
		}
	}
	if len(errs) != 0 {
		for _, err := range errs {
			fmt.Fprintln(os.Stderr, err)
		}
		os.Exit(1)
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
