package main

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/andrewarchi/browser/jsonutil"
)

func main() {
	var projects []Project
	try(jsonutil.DecodeFile("../projects.json", &projects))
	var b strings.Builder
	try(renderTable(&b, projects))
}

func try(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

type Project struct {
	Name        string   `json:"name"`
	Path        string   `json:"path"`
	Authors     []string `json:"authors"`
	Languages   []string `json:"languages"`
	Tags        []string `json:"tags"`
	Date        string   `json:"date"`
	SpecVersion string   `json:"spec_version"`
	Source      []string `json:"source"`
	Mirrored    bool     `json:"mirrored,omitempty"`
}

func renderTable(b *strings.Builder, projects []Project) error {
	for _, p := range projects {
		b.WriteString("| ")
		if p.Path != "" {
			renderLink(b, p.Name, p.Path)
		} else {
			b.WriteString(p.Name)
		}
		// TODO padding and additional fields
		b.WriteString(" |")
		for i, s := range p.Source {
			label, err := getURLLabel(s)
			if err != nil {
				return err
			}
			if i != 0 {
				b.WriteByte(',')
			}
			b.WriteByte(' ')
			renderLink(b, label, s)
		}
		b.WriteString(" |\n")
	}
	return nil
}

func renderLink(b *strings.Builder, label, url string) error {
	// TODO escape ] and )
	if strings.IndexByte(label, ']') != -1 {
		return fmt.Errorf("label contains ]: %s", label)
	}
	if strings.IndexByte(url, ')') != -1 {
		return fmt.Errorf("URL contains ): %s", url)
	}
	fmt.Fprintf(b, "[%s](%s)", label, url)
	return nil
}

var domainLabels = map[string]string{
	"github.com":                 "GitHub",
	"gitlab.com":                 "GitLab",
	"news.ycombinator.com":       "HN",
	"codegolf.stackexchange.com": "Code Golf",
	"code.activestate.com":       "ActiveState Code",
	"compsoc.dur.ac.uk":          "CompSoc",
	"cs.newcastle.edu.au":        "Newcastle",
}

func getURLLabel(rawURL string) (string, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}
	if u.Hostname() == "web.archive.org" && strings.HasPrefix(u.Path, "/web/") {
		path := strings.TrimPrefix(u.Path, "/web/")
		if i := strings.IndexByte(path, '/'); i != -1 {
			u, err = url.Parse(path[i+1:])
			if err != nil {
				return "", err
			}
		}
	}
	host := strings.TrimPrefix(u.Hostname(), "www.")
	if host == "compsoc.dur.ac.uk" && strings.HasPrefix(u.Path, "/archives/whitespace/") {
		return "Mailing list", nil
	}
	if label, ok := domainLabels[host]; ok {
		return label, nil
	}
	return host, nil
}
