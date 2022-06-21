package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"

	"github.com/wspace/corpus/tools"
	"golang.org/x/sys/execabs"
)

func main() {
	projects, err := tools.ReadAllProjects()
	try(err)
	tools.SortProjectsByID(projects)

	var badURLs []*tools.Project
	for _, p := range projects {
		if p.ID == "" || len(p.Source) == 0 {
			continue
		}
		repo, branch := getGitURL(p.Source[0])
		if repo == "" {
			badURLs = append(badURLs, p)
			continue
		}
		if _, err := os.Stat(p.ID); err == nil {
			continue
		}

		var cmd *exec.Cmd
		if branch != "" {
			fmt.Printf("git submodule add -b %s %s %s\n", branch, repo, p.ID)
			cmd = execabs.Command("git", "submodule", "add", "-b", branch, repo, p.ID)
		} else {
			fmt.Printf("git submodule add %s %s\n", repo, p.ID)
			cmd = execabs.Command("git", "submodule", "add", repo, p.ID)
		}
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		try(cmd.Run())
	}

	if len(badURLs) != 0 {
		fmt.Println("First source not a recognized repo for:")
		for _, p := range badURLs {
			url := p.Source[0]
			if label, err := tools.GetURLLabel(p.Source[0]); err == nil && label != "" {
				url = label
			}
			fmt.Printf("- %s: %s\n", p.ID, url)
		}
	}
}

var (
	github      = regexp.MustCompile(`^(https://github\.com/[^/]+/[^/]+)(?:/tree/([^/]+))?$`)
	githubGist  = regexp.MustCompile(`^https://gist\.github\.com/[^/]+/[^/]+$`)
	gitlab      = regexp.MustCompile(`^(https://gitlab\.com/[^/]+/[^/]+)(?:/-/tree/([^/]+))?$`)
	bitbucket   = regexp.MustCompile(`^(https://bitbucket\.org/[^/]+/[^/]+)(?:/src/([^/]+))?$`)
	sourceForge = regexp.MustCompile(`^https://git\.code\.sf\.net/p/[^/]+/code$`)
)

func getGitURL(url string) (gitURL, branch string) {
	if match := github.FindStringSubmatch(url); match != nil {
		return match[1], match[2]
	}
	if githubGist.MatchString(url) {
		return url, ""
	}
	if match := gitlab.FindStringSubmatch(url); match != nil {
		return match[1] + ".git", match[2]
	}
	if match := bitbucket.FindStringSubmatch(url); match != nil {
		return match[1], match[2]
	}
	if sourceForge.MatchString(url) {
		return url, ""
	}
	return "", ""
}

func try(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
