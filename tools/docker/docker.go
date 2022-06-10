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
		if p.Build != nil {
			dw.Reset()
			dw.Inst("FROM %s", p.Build.BaseImage)
			dw.Line("")
			dw.WorkDir("/home")
			dw.Inst("RUN git clone %s", p.Source[0])
			dir := "/home" + p.Source[0][strings.LastIndexByte(p.Source[0], '/'):]
			if len(p.Build.Setup) != 0 {
				dw.WorkDir(dir)
				for _, cmd := range p.Build.Setup {
					dw.RunShell(cmd)
				}
			}
			for _, target := range p.Build.Targets {
				dw.WorkDir(path.Join(dir, target.WorkDir))
				dw.RunShell(target.Build)
				dw.b.WriteString("# builds: ")
				for i, binary := range target.Binaries {
					if i != 0 {
						dw.b.WriteString(", ")
					}
					dw.b.WriteString(path.Join(dir, binary))
				}
				dw.b.WriteByte('\n')
			}
			try(dw.SaveIfChanged(p.ID + ".dockerfile"))
		}
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

func (dw *DockerfileWriter) Inst(format string, args ...any) {
	fmt.Fprintf(dw.b, format+"\n", args...)
}

func (dw *DockerfileWriter) Run(cmd string) {
	if strings.ContainsAny(cmd, `'"$*?><|&`) {
		dw.RunShell(cmd)
	} else {
		dw.RunExec(strings.Split(cmd, " "))
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

func (dw *DockerfileWriter) RunShell(cmd string) {
	dw.Inst("RUN %s", cmd)
}

func (dw *DockerfileWriter) WorkDir(dir string) {
	if dw.workDir != dir {
		dw.Inst("WORKDIR %s", dir)
		dw.workDir = dir
	}
}

func (dw *DockerfileWriter) Line(line string) {
	dw.b.WriteString(line + "\n")
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
