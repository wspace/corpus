package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/andrewarchi/browser/jsonutil"
	"github.com/wspace/corpus/tools"
)

func main() {
	var projects []tools.Project
	try(jsonutil.DecodeFile("projects.json", &projects))
	t, err := template.ParseFiles("README.md.tmpl")
	try(err)
	f, err := os.Create("README.md")
	try(err)
	var b strings.Builder
	try(tools.RenderProjectTable(&b, projects))
	try(t.Execute(f, map[string]interface{}{
		"projects": b.String(),
	}))
}

func try(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
