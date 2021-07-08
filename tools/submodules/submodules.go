package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"

	"github.com/andrewarchi/browser/jsonutil"
	"github.com/wspace/corpus/tools"
	"golang.org/x/sys/execabs"
)

func main() {
	var projects []tools.Project
	try(jsonutil.DecodeFile("projects.json", &projects))
	sort.Slice(projects, func(i, j int) bool {
		return projects[i].Path < projects[j].Path
	})

	for _, p := range projects {
		if p.Path == "" || len(p.Source) == 0 {
			continue
		}
		if _, err := os.Stat(p.Path); err == nil {
			continue
		}

		src := p.Source[0]
		repo := getGitURL(src)
		if repo == "" {
			fmt.Printf("%s: first source not a recognized repo: %s", p.Path, src)
			continue
		}

		fmt.Printf("git submodule add %s %s\n", repo, p.Path)
		cmd := execabs.Command("git", "submodule", "add", repo, p.Path)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		try(cmd.Run())
	}
}

var (
	github      = regexp.MustCompile(`^https://(?:gist\.)?github\.com/[^/]+/[^/]+$`)
	gitlab      = regexp.MustCompile(`^https://gitlab\.com/[^/]+/[^/]+$`)
	sourceforge = regexp.MustCompile(`^https://sourceforge\.net/projects/([^/]+)/`)
)

func getGitURL(url string) string {
	if github.MatchString(url) {
		return url
	} else if gitlab.MatchString(url) {
		return url + ".git"
	} else if m := sourceforge.FindStringSubmatch(url); m != nil {
		return "https://git.code.sf.net/p/" + m[1] + "/code"
	}
	return ""
}

func try(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
