package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/wspace/corpus/tools"
)

func main() {
	projects, err := tools.ReadAllProjects()
	try(err)
	tools.SortProjectsByTime(projects)
	t, err := template.ParseFiles("README.md.tmpl")
	try(err)
	f, err := os.Create("README.md")
	try(err)
	var b strings.Builder
	try(tools.RenderProjectTable(&b, projects))
	try(t.Execute(f, map[string]interface{}{
		"projects":      projects,
		"projectsTable": b.String(),
	}))
}

func try(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
