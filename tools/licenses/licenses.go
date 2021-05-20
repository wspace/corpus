package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/andrewarchi/browser/jsonutil"
	"github.com/andrewarchi/ws-corpus/tools"
)

var ghRepo = regexp.MustCompile("^https://github.com/[^/]+/[^/]+$")
var ghToken = os.Getenv("GITHUB_ACCESS_TOKEN")

func main() {
	var projects []tools.Project
	f, err := os.OpenFile("projects.json", os.O_RDWR, 0o644)
	try(err)
	try(jsonutil.Decode(f, &projects))
	for i := range projects {
		p := &projects[i]
		if p.License != "" || len(p.Source) == 0 || !ghRepo.MatchString(p.Source[0]) {
			continue
		}
		repo := p.Source[0]
		fmt.Printf("Getting license for %s\n", repo)
		repo = strings.TrimPrefix(repo, "https://github.com/")
		license, err := getGitHubLicense(repo)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			break
		}
		p.License = license
	}
	try(f.Truncate(0))
	_, err = f.Seek(0, io.SeekStart)
	try(err)
	e := json.NewEncoder(f)
	e.SetEscapeHTML(false)
	try(e.Encode(projects))
}

func getGitHubLicense(repo string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://api.github.com/repos/%s/license", repo), nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("Accept", "application/vnd.github.v3+json")
	if ghToken != "" {
		req.Header.Add("Authorization", "token "+ghToken)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var l struct {
		Message string `json:"message"`
		License struct {
			ID string `json:"spdx_id"`
		} `json:"license"`
	}
	if err := jsonutil.DecodeAllowUnknownFields(resp.Body, &l); err != nil {
		return "", err
	}
	if l.License.ID != "" {
		if l.License.ID == "NOASSERTION" {
			return "other", nil
		}
		return l.License.ID, nil
	}
	if l.Message != "" {
		if l.Message == "Not Found" {
			return "not found", nil
		}
		return "", fmt.Errorf("message: %s", l.Message)
	}
	return "", nil
}

func try(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
