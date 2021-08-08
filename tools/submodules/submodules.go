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

	var badURLs []*tools.Project
	for i := range projects {
		p := &projects[i]
		if p.Path == "" || len(p.Source) == 0 {
			continue
		}
		if _, err := os.Stat(p.Path); err == nil {
			continue
		}

		repo := getGitURL(p.Source[0])
		if repo == "" {
			badURLs = append(badURLs, p)
			continue
		}

		fmt.Printf("git submodule add %s %s\n", repo, p.Path)
		cmd := execabs.Command("git", "submodule", "add", repo, p.Path)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		try(cmd.Run())
	}

	if len(badURLs) != 0 {
		fmt.Println("First source not a recognized repo for:")
		for _, p := range badURLs {
			fmt.Printf("- %s: %s\n", p.Path, p.Source[0])
		}
	}
}

var (
	github      = regexp.MustCompile(`^https://(?:gist\.)?github\.com/[^/]+/[^/]+$`)
	gitlab      = regexp.MustCompile(`^https://gitlab\.com/[^/]+/[^/]+$`)
	bitbucket   = regexp.MustCompile(`^https://bitbucket\.org/[^/]+/[^/]+$`)
	sourceforge = regexp.MustCompile(`^https://sourceforge\.net/projects/([^/]+)/`)
)

func getGitURL(url string) string {
	if github.MatchString(url) || bitbucket.MatchString(url) {
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
