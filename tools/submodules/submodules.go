package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/wspace/corpus/tools"
	"golang.org/x/sys/execabs"
)

func main() {
	verbose := len(os.Args) > 1 && os.Args[1] == "-v"
	projects, err := tools.ReadAllProjects()
	try(err)
	tools.SortProjectsByID(projects)

	var badURLs []*tools.Project
	var noRepoName []*tools.Project
	for _, p := range projects {
		if p.ID == "" || len(p.Source) == 0 {
			continue
		}
		repo, branch, _ := tools.GetGitURL(p.Source[0])
		if repo == "" {
			badURLs = append(badURLs, p)
			continue
		}
		repoName := p.RepoName()
		if repoName == "" {
			noRepoName = append(noRepoName, p)
			continue
		}
		submodule := p.ID + "/" + repoName
		if _, err := os.Stat(submodule); err == nil {
			continue
		}

		var cmd *exec.Cmd
		if branch != "" {
			fmt.Printf("git submodule add -b %s %s %s\n", branch, repo, submodule)
			cmd = execabs.Command("git", "submodule", "add", "-b", branch, repo, submodule)
		} else {
			fmt.Printf("git submodule add %s %s\n", repo, submodule)
			cmd = execabs.Command("git", "submodule", "add", repo, submodule)
		}
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		try(cmd.Run())
	}

	listProjects := func(projects []*tools.Project, problem string) {
		if len(projects) != 0 {
			if verbose {
				fmt.Printf("%s for:\n", problem)
				for _, p := range projects {
					url := p.Source[0]
					if label, err := tools.GetURLLabel(p.Source[0]); err == nil && label != "" {
						url = label
					}
					fmt.Printf("- %s: %s\n", p.ID, url)
				}
			} else {
				fmt.Printf("%s for %d projects\n", problem, len(projects))
			}
		}
	}
	listProjects(badURLs, "First source not a recognized repo")
	listProjects(noRepoName, "Repo name could not be determined")
}

func try(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
