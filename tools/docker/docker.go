package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"
	"strconv"
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
	dw := NewDockerfileWriter()
	for _, p := range projects {
		for _, cmd := range p.Commands {
			if cmd.Build != "" || len(cmd.Dependencies) != 0 || cmd.InstallDependencies != "" {
				goto has_build
			}
		}
		continue
	has_build:

		dw.Reset()
		dw.Write("FROM alpine AS builder")
		dw.Write("")

		dir := p.RepoName()
		if dir == "" {
			dir = "TODO"
		}
		dw.WorkDir("/" + dir)
		dw.Write("COPY %s .", dir)
		dir = "/" + dir
		for _, cmd := range p.Commands {
			dw.Write("")
			if cmd.Bin != "" {
				dw.Write("# Build %s", path.Join(dir, cmd.Bin))
			}
			if len(cmd.Dependencies) != 0 {
				dw.Write("# Dependencies: %s", strings.Join(cmd.Dependencies, ", "))
			}
			dw.RunAnd(cmd.InstallDependencies)
			if cmd.BuildErrors != "" {
				dw.Write("# Build errors: %s", cmd.BuildErrors)
			}
			dw.RunAnd(cmd.Build)
		}

		if len(p.Commands) != 0 {
			dw.Write("")
			dw.Write("FROM scratch")
			dw.Write("")
			dw.Write("WORKDIR %s", dir)
			var entrypoint string
			for _, cmd := range p.Commands {
				if cmd.Bin != "" {
					dw.Write("COPY --from=builder %s/%s .", dir, cmd.Bin)
					if entrypoint == "" {
						if cmd.Usage != nil {
							entrypoint = `"` + strings.ReplaceAll(strings.ReplaceAll(*cmd.Usage, "$0", cmd.Bin), " ", `", "`) + `"`
						} else {
							entrypoint = strconv.Quote(dir + "/" + path.Base(cmd.Bin))
						}
					}
				}
			}
			if entrypoint != "" {
				dw.Write(`ENTRYPOINT [%s]`, entrypoint)
			}
		}

		try(dw.SaveIfChanged(p.ID + "/Dockerfile"))
	}
	if len(errs) != 0 {
		for _, err := range errs {
			fmt.Fprintln(os.Stderr, err)
		}
		os.Exit(1)
	}
}

type DockerfileWriter struct {
	b       *bytes.Buffer
	workDir string
}

func NewDockerfileWriter() *DockerfileWriter {
	return &DockerfileWriter{b: new(bytes.Buffer)}
}

func (dw *DockerfileWriter) Write(format string, args ...any) {
	fmt.Fprintf(dw.b, format+"\n", args...)
}

func (dw *DockerfileWriter) Run(format string, args ...any) {
	dw.Write("RUN "+format, args...)
}

func (dw *DockerfileWriter) RunAnd(cmd string) {
	if cmd != "" {
		for _, cmd := range strings.Split(cmd, " && ") {
			dw.Run(cmd)
		}
	}
}

func (dw *DockerfileWriter) RunExec(cmd []string) {
	dw.b.WriteString("RUN [")
	for i, s := range cmd {
		if i != 0 {
			dw.b.WriteString(", ")
		}
		dw.b.WriteString(strconv.Quote(s))
	}
	dw.b.WriteString("]\n")
}

func (dw *DockerfileWriter) WorkDir(dir string) {
	if dw.workDir != dir {
		dw.Write("WORKDIR %s", dir)
		dw.workDir = dir
	}
}

func (dw *DockerfileWriter) Reset() {
	dw.b.Reset()
	dw.workDir = ""
}

func (dw *DockerfileWriter) SaveIfChanged(filename string) error {
	contents, err := os.ReadFile(filename)
	if (err == nil && !bytes.Equal(contents, dw.b.Bytes())) || errors.Is(err, fs.ErrNotExist) {
		fmt.Printf("Generated %s\n", filename)
		return os.WriteFile(filename, dw.b.Bytes(), 0o644)
	}
	return err
}

func try(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
