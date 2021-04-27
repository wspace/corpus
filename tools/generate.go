package main

import (
	"fmt"
	"net/url"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/andrewarchi/browser/jsonutil"
)

func main() {
	var projects []Project
	try(jsonutil.DecodeFile("../projects.json", &projects))
	var b strings.Builder
	try(renderTable(&b, projects))
	fmt.Println(b.String())
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
	Features    struct {
		ArbitraryPrecision bool              `json:"arbitrary_precision"`
		Assembly           map[string]string `json:"assembly"`
	} `json:"features,omitempty"`
}

func renderTable(b *strings.Builder, projects []Project) error {
	padding := []int{46, 16, 11, 12, 10, 4, 0}
	renderColumns(b, padding, "Name", "Author(s)", "Language(s)", "Tags", "Date", "Spec", "Source")
	b.WriteString("| ---------------------------------------------- | ---------------- | ----------- | ------------ | ---------- | ---- | ------ |\n")
	for _, p := range projects {
		name := p.Name
		if p.Path != "" {
			var err error
			name, err = formatLink(p.Name, p.Path)
			if err != nil {
				return err
			}
		}
		links := make([]string, 0, len(p.Source))
		for _, s := range p.Source {
			label, err := getURLLabel(s)
			if err != nil {
				return err
			}
			link, err := formatLink(label, s)
			if err != nil {
				return err
			}
			links = append(links, link)
		}
		renderColumns(b, padding,
			name,
			strings.Join(p.Authors, ", "),
			strings.Join(p.Languages, ", "),
			strings.Join(p.Tags, ", "),
			p.Date,
			p.SpecVersion,
			strings.Join(links, ", "))
	}
	return nil
}

func renderColumns(b *strings.Builder, padding []int, cols ...string) {
	width := 0
	padWidth := 0
	for i, col := range cols {
		b.WriteString("| ")
		b.WriteString(col)
		width += utf8.RuneCountInString(col)
		if i < len(padding) && padding[i] != 0 {
			padWidth += padding[i]
			for ; width < padWidth; width++ {
				b.WriteByte(' ')
			}
			b.WriteByte(' ')
		} else if len(col) != 0 {
			b.WriteByte(' ')
		}
	}
	b.WriteString("|\n")
}

func formatLink(label, url string) (string, error) {
	// TODO escape ] and )
	if strings.IndexByte(label, ']') != -1 {
		return "", fmt.Errorf("label contains ]: %s", label)
	}
	if strings.IndexByte(url, ')') != -1 {
		return "", fmt.Errorf("URL contains ): %s", url)
	}
	return fmt.Sprintf("[%s](%s)", label, url), nil
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
