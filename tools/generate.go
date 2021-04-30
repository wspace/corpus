package main

import (
	"fmt"
	"net/url"
	"os"
	"regexp"
	"strings"
	"text/template"
	"time"
	"unicode/utf8"

	"github.com/andrewarchi/browser/jsonutil"
)

func main() {
	var projects []Project
	try(jsonutil.DecodeFile("projects.json", &projects))
	t, err := template.ParseFiles("README.md.tmpl")
	try(err)
	f, err := os.Create("README.md")
	try(err)
	var b strings.Builder
	try(renderTable(&b, projects))
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

type Project struct {
	Name        string   `json:"name"`
	Path        string   `json:"path"`
	Authors     []string `json:"authors"`
	Languages   []string `json:"languages"`
	Tags        []string `json:"tags"`
	Date        string   `json:"date"`
	SpecVersion string   `json:"spec_version"`
	Source      []string `json:"source"`
	Features    struct {
		ArbitraryPrecision bool `json:"arbitrary_precision,omitempty"`
	} `json:"features,omitempty"`
	Assembly *struct {
		Instructions              map[string][]string `json:"instructions,omitempty"` // key: instruction name, aliases
		CaseSensitiveInstructions bool                `json:"case_sensitive_instructions,omitempty"`
		LineCommentPrefix         string              `json:"line_comment_prefix,omitempty"`
		LabelDefFormat            string              `json:"label_def_format,omitempty"`
		LabelRefFormat            string              `json:"label_ref_format,omitempty"`
		Extension                 string              `json:"extension,omitempty"`
	} `json:"assembly,omitempty"`
	Mapping *struct {
		Space         string `json:"space"`
		Tab           string `json:"tab"`
		LF            string `json:"lf"`
		SpacesBetween string `json:"spaces_between,omitempty"`
	} `json:"mapping,omitempty"`
	Notes string `json:"notes,omitempty"`
}

func renderTable(b *strings.Builder, projects []Project) error {
	padding := []int{46, 16, 10, 12, 10, 3, 0}
	head := []string{"Name", "Authors", "Languages", "Tags", "Date", "Spec", "Source"}
	renderRow(b, padding, head, false)
	b.WriteByte('\n')
	renderRow(b, padding, head, true)
	for _, p := range projects {
		b.WriteByte('\n')
		row, err := p.formatColumns()
		if err != nil {
			return err
		}
		renderRow(b, padding, row, false)
	}
	return nil
}

func (p *Project) formatColumns() ([]string, error) {
	name := p.Name
	if p.Path != "" {
		name = formatLink(p.Name, p.Path)
	}
	date := p.Date
	if t, err := time.Parse("2006-01-02 15:04:05 -0700", date); err == nil {
		date = t.Format("2006-01-02")
	}
	links := make([]string, 0, len(p.Source))
	for _, s := range p.Source {
		label, err := getURLLabel(s)
		if err != nil {
			return nil, err
		}
		links = append(links, formatLink(label, s))
	}
	return []string{
		name,
		strings.Join(p.Authors, ", "),
		strings.Join(p.Languages, ", "),
		strings.Join(p.Tags, ", "),
		date,
		p.SpecVersion,
		strings.Join(links, ", ")}, nil
}

func renderRow(b *strings.Builder, padding []int, row []string, dashes bool) {
	if len(padding) != len(row) {
		panic("row length mismatch")
	}
	width := 0
	padWidth := 0
	pad := byte(' ')
	if dashes {
		pad = '-'
	}
	for i, col := range row {
		b.WriteString("| ")
		n := utf8.RuneCountInString(col)
		width += n
		padWidth += padding[i]
		if dashes {
			for i := 0; i < n; i++ {
				b.WriteByte('-')
			}
		} else {
			b.WriteString(col)
		}
		if padding[i] != 0 {
			for ; width < padWidth; width++ {
				b.WriteByte(pad)
			}
		}
		if col != "" || padding[i] != 0 {
			b.WriteByte(' ')
		}
	}
	b.WriteByte('|')
}

func formatLink(label, url string) string {
	label = strings.ReplaceAll(label, "]", `\]`)
	url = strings.ReplaceAll(url, ")", `\)`)
	return fmt.Sprintf("[%s](%s)", label, url)
}

var domainLabels = map[string]string{
	"github.com":                 "GitHub",
	"gitlab.com":                 "GitLab",
	"news.ycombinator.com":       "HN",
	"codegolf.stackexchange.com": "Code Golf",
	"code.activestate.com":       "ActiveState Code",
	"compsoc.dur.ac.uk":          "CompSoc",
	"cs.newcastle.edu.au":        "Newcastle",
	"what.thedailywtf.com":       "What the Daily WTF?",
}

var subSites = map[string]struct{}{
	"blogspot.com": {},
}

var pathPatterns = map[string]*regexp.Regexp{
	"reddit.com": regexp.MustCompile("/(r/[^/]+).*"),
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
	if i := strings.IndexByte(host, '.'); i != -1 {
		if _, ok := subSites[host[i+1:]]; ok {
			return host[:i], nil
		}
	}
	if pattern, ok := pathPatterns[host]; ok {
		if label := pattern.ReplaceAllString(u.Path, "$1"); label != "" {
			return label, nil
		}
	}
	if label, ok := domainLabels[host]; ok {
		return label, nil
	}
	return host, nil
}
