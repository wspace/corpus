package tools

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/andrewarchi/browser/jsonutil"
	"github.com/pelletier/go-toml"
)

// Uses SPDX license IDs
// https://spdx.org/licenses/

// Stack Exchange content licenses
// https://meta.stackexchange.com/help/licensing
// CC BY-SA 2.5 before 2011-04-08
// CC BY-SA 3.0 from 2011-04-08 up to but not including 2018-05-02
// CC BY-SA 4.0 on or after 2018-05-02

var ghRepo = regexp.MustCompile("^https://github.com/[^/]+/[^/]+$")
var ghToken = os.Getenv("GITHUB_ACCESS_TOKEN")

func (p *Project) GetLicense() (string, error) {
	if len(p.Source) > 0 && ghRepo.MatchString(p.Source[0]) {
		fmt.Fprintf(os.Stderr, "Getting license for %s from GitHub\n", p.ID)
		repo := strings.TrimPrefix(p.Source[0], "https://github.com/")
		if l, err := getGitHubLicense(repo); l != "" || err != nil {
			return l, err
		}
	}
	if repoName := p.RepoName(); repoName != "" && p.ID != "" {
		submodule := p.ID + "/" + repoName
		if l := getLicenseFromFile(submodule, "package.json", getPackageJSONLicense); l != "" {
			return l, nil
		} else if l := getLicenseFromFile(submodule, "Cargo.toml", getCargoTOMLLicense); l != "" {
			return l, nil
		}
	}
	return "", nil
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
	if l.Message != "" && l.Message != "Not Found" {
		return "", fmt.Errorf("message: %s", l.Message)
	}
	return "", nil
}

func getLicenseFromFile(path, filename string, f func(string) (string, error)) string {
	filename = path + "/" + filename
	if stat, err := os.Stat(filename); err != nil || stat.IsDir() {
		return ""
	}
	l, err := f(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	if l != "" {
		fmt.Fprintf(os.Stderr, "Got license for %s from %s\n", path, filename)
	}
	return l
}

func getPackageJSONLicense(filename string) (string, error) {
	var pack struct {
		License string `json:"license"`
	}
	if err := jsonutil.DecodeFileAllowUnknownFields(filename, &pack); err != nil {
		return "", err
	}
	return pack.License, nil
}

func getCargoTOMLLicense(filename string) (string, error) {
	tree, err := toml.LoadFile(filename)
	if err != nil {
		return "", err
	}
	if l, ok := tree.Get("package.license").(string); ok {
		return l, nil
	}
	return "", nil
}
