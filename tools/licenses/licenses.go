package main

import (
	"fmt"
	"os"

	"github.com/wspace/corpus/tools"
)

func main() {
	projects, err := tools.ReadProjects(os.Args[1:])
	try(err)
	for _, p := range projects {
		if p.License == "" && p.ID != "" {
			if l, err := p.GetLicense(); err != nil {
				fmt.Fprintln(os.Stderr, err)
				break
			} else if l != "" {
				p.License = l
			} else {
				p.License = "not found"
			}
			p.Write()
		}
	}
}

func try(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
