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
	Description string   `json:"description"`
	Path        string   `json:"path"`
	Authors     []string `json:"authors"`
	Languages   []string `json:"languages"`
	Tags        []string `json:"tags"`
	Date        string   `json:"date"`
	SpecVersion string   `json:"spec_version"`
	Source      []string `json:"source"`
	Bounds      struct {
		Precision      string      `json:"precision,omitempty"`
		LabelPrecision string      `json:"label_precision,omitempty"` // usually the same as precision
		StackCap       interface{} `json:"stack_cap,omitempty"`       // int or string
		CallStackCap   interface{} `json:"call_stack_cap,omitempty"`  // int or string
		HeapMin        interface{} `json:"heap_min,omitempty"`        // int or string
		HeapMax        interface{} `json:"heap_max,omitempty"`        // int or string
		HeapCap        interface{} `json:"heap_cap,omitempty"`        // int or string
		LabelCap       interface{} `json:"label_cap,omitempty"`       // int or string
		InstructionCap interface{} `json:"instruction_cap,omitempty"` // int or string
	} `json:"bounds,omitempty"`
	Behavior struct {
		BufferedOutput bool   `json:"buffered_output,omitempty"`
		EOF            string `json:"eof,omitempty"`
	} `json:"behavior,omitempty"`
	Assembly *struct {
		Instructions              map[Instruction][]string `json:"instructions,omitempty"`
		CaseSensitiveInstructions bool                     `json:"case_sensitive_instructions,omitempty"`
		LineCommentPrefix         string                   `json:"line_comment_prefix,omitempty"`
		LabelDefFormat            string                   `json:"label_def_format,omitempty"`
		LabelRefFormat            string                   `json:"label_ref_format,omitempty"`
		Extension                 string                   `json:"extension,omitempty"`
	} `json:"assembly,omitempty"`
	Mapping *struct {
		Space         string `json:"space"`
		Tab           string `json:"tab"`
		LF            string `json:"lf"`
		SpacesBetween bool   `json:"spaces_between,omitempty"`
		IgnoreCase    bool   `json:"ignore_case,omitempty"`
		Extension     string `json:"extension,omitempty"`
	} `json:"mapping,omitempty"`
	Commands []Command `json:"commands,omitempty"`
	Notes    string    `json:"notes,omitempty"`
}

type Command struct {
	Type                string   `json:"type"`
	Bin                 string   `json:"bin"`
	InstallDependencies []string `json:"install_dependencies,omitempty"`
	Build               string   `json:"build,omitempty"`
	BuildErrors         string   `json:"build_errors,omitempty"`
	Usage               string   `json:"usage,omitempty"`
	Output              string   `json:"output,omitempty"`
	Options             []struct {
		Short       string      `json:"short,omitempty"` // -s
		Long        string      `json:"long,omitempty"`  // --long
		Arg         string      `json:"arg,omitempty"`
		ArgRequired bool        `json:"arg_required,omitempty"`
		Default     interface{} `json:"default,omitempty"`
		Desc        string      `json:"desc,omitempty"`
		Values      []struct {
			Values []string `json:"values,omitempty"`
			Desc   string   `json:"desc,omitempty"`
		} `json:"values,omitempty"`
	} `json:"options,omitempty"`
	OptionParse string `json:"option_parse,omitempty"`
	Subcommands []struct {
		Name string `json:"name"`
		Desc string `json:"desc,omitempty"`
	} `json:"subcommands,omitempty"`
	Notes string `json:"notes,omitempty"`
}

type Instruction uint8

const (
	Push Instruction = iota + 1
	Dup
	Copy
	Swap
	Drop
	Slide
	Add
	Sub
	Mul
	Div
	Mod
	Store
	Retrieve
	Label
	Call
	Jmp
	Jz
	Jn
	Ret
	End
	Printc
	Printi
	Readc
	Readi
)

func (inst *Instruction) UnmarshalText(text []byte) error {
	switch string(text) {
	case "push":
		*inst = Push
	case "dup":
		*inst = Dup
	case "copy":
		*inst = Copy
	case "swap":
		*inst = Swap
	case "drop":
		*inst = Drop
	case "slide":
		*inst = Slide
	case "add":
		*inst = Add
	case "sub":
		*inst = Sub
	case "mul":
		*inst = Mul
	case "div":
		*inst = Div
	case "mod":
		*inst = Mod
	case "store":
		*inst = Store
	case "retrieve":
		*inst = Retrieve
	case "label":
		*inst = Label
	case "call":
		*inst = Call
	case "jmp":
		*inst = Jmp
	case "jz":
		*inst = Jz
	case "jn":
		*inst = Jn
	case "ret":
		*inst = Ret
	case "end":
		*inst = End
	case "printc":
		*inst = Printc
	case "printi":
		*inst = Printi
	case "readc":
		*inst = Readc
	case "readi":
		*inst = Readi
	default:
		return fmt.Errorf("illegal instruction: %s", text)
	}
	return nil
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
	"gist.github.com":            "GitHub Gist",
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
			label, err := getURLLabel(path[i+1:])
			if err != nil {
				return "", err
			}
			return label + " (archive)", nil
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
