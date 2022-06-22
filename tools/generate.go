package tools

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/andrewarchi/browser/jsonutil"
	"golang.org/x/sys/execabs"
)

type Project struct {
	ID                string   `json:"id"`
	Name              string   `json:"name"`
	Description       string   `json:"description,omitempty"`
	Authors           []string `json:"authors"`
	License           string   `json:"license"`
	Languages         []string `json:"languages"`
	Tags              []string `json:"tags"`
	Date              string   `json:"date"`
	SpecVersion       string   `json:"spec_version"`
	Source            []string `json:"source"`
	SourceUnavailable bool     `json:"source_unavailable,omitempty"`
	Package           *struct {
		Name    string `json:"name"`
		Manager string `json:"manager"`
		URL     string `json:"url"`
	} `json:"package,omitempty"`
	Relations []struct {
		ID      string `json:"id"`
		Type    string `json:"type"`
		Commit  string `json:"commit,omitempty"`
		Release string `json:"release,omitempty"`
	} `json:"relations,omitempty"`
	Challenges []string `json:"challenges,omitempty"`
	Bounds     *struct {
		Precision      string      `json:"precision,omitempty"`
		ArgPrecision   string      `json:"arg_precision,omitempty"`   // usually the same as precision
		LabelPrecision string      `json:"label_precision,omitempty"` // usually the same as precision
		StackCap       interface{} `json:"stack_cap,omitempty"`       // int or string
		CallStackCap   interface{} `json:"call_stack_cap,omitempty"`  // int or string
		HeapMin        interface{} `json:"heap_min,omitempty"`        // int or string
		HeapMax        interface{} `json:"heap_max,omitempty"`        // int or string
		HeapCap        interface{} `json:"heap_cap,omitempty"`        // int or string
		LabelCap       interface{} `json:"label_cap,omitempty"`       // int or string
		InstructionCap interface{} `json:"instruction_cap,omitempty"` // int or string
	} `json:"bounds,omitempty"`
	Behavior *struct {
		BufferedOutput *bool  `json:"buffered_output,omitempty"`
		EOF            string `json:"eof,omitempty"`
	} `json:"behavior,omitempty"`
	Whitespace *struct {
		Unimplemented []string `json:"unimplemented,omitempty"`
		NonStandard   []struct {
			Name    string   `json:"name,omitempty"`
			Aliases []string `json:"aliases,omitempty"`
			Seq     string   `json:"seq,omitempty"`
			Args    []string `json:"args,omitempty"`
			Desc    string   `json:"desc,omitempty"`
		} `json:"nonstandard,omitempty"`
		Extension string `json:"extension,omitempty"`
	} `json:"whitespace,omitempty"`
	Assembly *struct {
		Mnemonics *struct {
			Push      []string `json:"push,omitempty"`
			Dup       []string `json:"dup,omitempty"`
			Copy      []string `json:"copy,omitempty"`
			Swap      []string `json:"swap,omitempty"`
			Drop      []string `json:"drop,omitempty"`
			Slide     []string `json:"slide,omitempty"`
			Add       []string `json:"add,omitempty"`
			Sub       []string `json:"sub,omitempty"`
			Mul       []string `json:"mul,omitempty"`
			Div       []string `json:"div,omitempty"`
			Mod       []string `json:"mod,omitempty"`
			Store     []string `json:"store,omitempty"`
			Retrieve  []string `json:"retrieve,omitempty"`
			Label     []string `json:"label,omitempty"`
			Call      []string `json:"call,omitempty"`
			Jmp       []string `json:"jmp,omitempty"`
			Jz        []string `json:"jz,omitempty"`
			Jn        []string `json:"jn,omitempty"`
			Ret       []string `json:"ret,omitempty"`
			End       []string `json:"end,omitempty"`
			Printc    []string `json:"printc,omitempty"`
			Printi    []string `json:"printi,omitempty"`
			Readc     []string `json:"readc,omitempty"`
			Readi     []string `json:"readi,omitempty"`
			Shuffle   []string `json:"shuffle,omitempty"`
			DumpStack []string `json:"dumpstack,omitempty"`
			DumpHeap  []string `json:"dumpheap,omitempty"`
			DumpTrace []string `json:"dumptrace,omitempty"`
		} `json:"mnemonics,omitempty"`
		Macros []struct {
			Name    string   `json:"name"`
			Args    []string `json:"args,omitempty"`
			Replace string   `json:"replace,omitempty"`
			Notes   string   `json:"notes,omitempty"`
		} `json:"macros,omitempty"`
		Patterns                  map[string]string `json:"patterns,omitempty"`
		CaseSensitiveInstructions *bool             `json:"case_sensitive_instructions,omitempty"`
		MultiplePerLine           *bool             `json:"multiple_per_line,omitempty"`
		LineComments              []string          `json:"line_comments,omitempty"`
		BlockComments             []struct {
			Start    string `json:"start"`
			End      string `json:"end"`
			Nestable bool   `json:"nestable"`
		} `json:"block_comments,omitempty"`
		Indentation      *string  `json:"indentation,omitempty"`
		LabelIndentation *string  `json:"label_indentation,omitempty"`
		BlockIndentation *bool    `json:"block_indentation,omitempty"`
		BinaryNumbers    *bool    `json:"binary_numbers,omitempty"`
		Usage            []string `json:"usage,omitempty"`
		Extension        *string  `json:"extension,omitempty"`
		Notes            string   `json:"notes,omitempty"`
	} `json:"assembly,omitempty"`
	Mappings []struct {
		Space          string `json:"space"`
		Tab            string `json:"tab"`
		LF             string `json:"lf"`
		SpacesBetween  *bool  `json:"spaces_between,omitempty"`
		SpaceBeforeArg *bool  `json:"space_before_arg,omitempty"`
		LineComment    string `json:"line_comment,omitempty"`
		BeforeSTL      *bool  `json:"before_stl,omitempty"`
		IgnoreCase     *bool  `json:"ignore_case,omitempty"`
		Extension      string `json:"extension,omitempty"`
		Notes          string `json:"notes,omitempty"`
	} `json:"mappings,omitempty"`
	Programs []struct {
		Path         string   `json:"path"`
		Generated    string   `json:"generated,omitempty"`
		Inputs       []string `json:"inputs,omitempty"`
		Outputs      []string `json:"outputs,omitempty"`
		Aux          []string `json:"aux,omitempty"`
		Polyglot     []string `json:"polyglot,omitempty"`
		MappingIndex *int     `json:"mapping_index,omitempty"`
		Equivalent   string   `json:"equivalent,omitempty"`
		SpecVersion  string   `json:"spec_version,omitempty"`
		Generate     string   `json:"generate,omitempty"`
		Authors      []string `json:"authors,omitempty"`
		Desc         string   `json:"desc,omitempty"`
		Notes        string   `json:"notes,omitempty"`
	} `json:"programs,omitempty"`
	BuildErrors string    `json:"build_errors,omitempty"`
	RunErrors   string    `json:"run_errors,omitempty"`
	Commands    []Command `json:"commands,omitempty"`
	Notes       string    `json:"notes,omitempty"`
	TODO        string    `json:"todo,omitempty"`
}

type Command struct {
	Type                string          `json:"type,omitempty"`
	Bin                 string          `json:"bin,omitempty"`
	Dependencies        []string        `json:"dependencies,omitempty"`
	InstallDependencies string          `json:"install_dependencies,omitempty"`
	Build               string          `json:"build,omitempty"`
	BuildErrors         string          `json:"build_errors,omitempty"`
	Usage               *string         `json:"usage,omitempty"`
	ExampleUsages       []string        `json:"example_usages,omitempty"`
	RunErrors           string          `json:"run_errors,omitempty"`
	Input               string          `json:"input,omitempty"`
	Output              string          `json:"output,omitempty"`
	Options             []CommandOption `json:"options,omitempty"`
	OptionParse         string          `json:"option_parse,omitempty"`
	Subcommands         []struct {
		Name    string          `json:"name"`
		Desc    string          `json:"desc,omitempty"`
		Usage   string          `json:"usage,omitempty"`
		Input   string          `json:"input,omitempty"`
		Output  string          `json:"output,omitempty"`
		Options []CommandOption `json:"options,omitempty"`
		Notes   string          `json:"notes,omitempty"`
	} `json:"subcommands,omitempty"`
	Unimplemented []string `json:"unimplemented,omitempty"`
	Notes         string   `json:"notes,omitempty"`
}

type CommandOption struct {
	Short         interface{} `json:"short,omitempty"` // -s
	Long          string      `json:"long,omitempty"`  // --long
	Bare          string      `json:"bare,omitempty"`  // bare
	Required      *bool       `json:"required,omitempty"`
	RepeatAllowed *bool       `json:"repeat_allowed,omitempty"`
	Arg           string      `json:"arg,omitempty"`
	ArgRequired   *bool       `json:"arg_required,omitempty"`
	Type          string      `json:"type,omitempty"`
	Default       interface{} `json:"default,omitempty"`
	Min           interface{} `json:"min,omitempty"`
	Desc          string      `json:"desc,omitempty"`
	Values        []struct {
		Values []string `json:"values,omitempty"`
		Desc   string   `json:"desc,omitempty"`
	} `json:"values,omitempty"`
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

	Shuffle
	DumpStack
	DumpHeap
	DumpTrace
)

func ReadAllProjects() ([]*Project, error) {
	paths, err := filepath.Glob("*/*/project.json")
	if err != nil {
		return nil, err
	}
	return ReadProjects(paths)
}

var projectRe = regexp.MustCompile(`/(?:project\.json|Dockerfile)?$`)

func ReadProjects(paths []string) ([]*Project, error) {
	var projects []*Project
	for _, project := range paths {
		if strings.HasPrefix(project, "tools/") || project[0] == '.' {
			continue
		}
		id := projectRe.ReplaceAllLiteralString(project, "")
		var p Project
		if err := jsonutil.DecodeFile(id+"/project.json", &p); err != nil {
			return projects, fmt.Errorf("%s: %w", project, err)
		}
		if p.ID != "" && p.ID != id {
			return projects, fmt.Errorf("%s: ID %q does not match path", project, p.ID)
		}
		projects = append(projects, &p)
	}
	return projects, nil
}

func (p *Project) Write() error {
	var b bytes.Buffer
	e := json.NewEncoder(&b)
	e.SetEscapeHTML(false)
	if err := e.Encode(p); err != nil {
		return err
	}
	cmd := execabs.Command("underscore", "print", "-o", p.ID+"/project.json")
	cmd.Stdin = &b
	return cmd.Run()
}

func SortProjectsByTime(projects []*Project) {
	sort.Slice(projects, func(i, j int) bool {
		pi, pj := projects[i], projects[j]
		ti, tj := pi.Time(), pj.Time()
		return ti.After(tj) || (ti.Equal(tj) && pi.ID < pj.ID)
	})
}

func SortProjectsByID(projects []*Project) {
	sort.Slice(projects, func(i, j int) bool { return projects[i].ID < projects[j].ID })
}

func (p *Project) Time() time.Time {
	for _, layout := range []string{
		"2006-01-02 15:04:05 -0700",
		"2006-01-02 15:04:05",
		"2006-01-02 15:04",
		"2006-01-02",
	} {
		if t, err := time.ParseInLocation(layout, p.Date, time.UTC); err == nil {
			return t
		}
	}
	return time.Time{}
}

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
	case "shuffle":
		*inst = Shuffle
	case "dumpstack":
		*inst = DumpStack
	case "dumpheap":
		*inst = DumpHeap
	case "dumptrace":
		*inst = DumpTrace
	default:
		return fmt.Errorf("illegal instruction: %s", text)
	}
	return nil
}

func (inst Instruction) String() string {
	switch inst {
	case Push:
		return "push"
	case Dup:
		return "dup"
	case Copy:
		return "copy"
	case Swap:
		return "swap"
	case Drop:
		return "drop"
	case Slide:
		return "slide"
	case Add:
		return "add"
	case Sub:
		return "sub"
	case Mul:
		return "mul"
	case Div:
		return "div"
	case Mod:
		return "mod"
	case Store:
		return "store"
	case Retrieve:
		return "retrieve"
	case Label:
		return "label"
	case Call:
		return "call"
	case Jmp:
		return "jmp"
	case Jz:
		return "jz"
	case Jn:
		return "jn"
	case Ret:
		return "ret"
	case End:
		return "end"
	case Printc:
		return "printc"
	case Printi:
		return "printi"
	case Readc:
		return "readc"
	case Readi:
		return "readi"
	case Shuffle:
		return "shuffle"
	case DumpStack:
		return "dumpstack"
	case DumpHeap:
		return "dumpheap"
	case DumpTrace:
		return "dumptrace"
	default:
		return fmt.Sprintf("instruction(%d)", uint8(inst))
	}
}

func RenderProjectTable(b *strings.Builder, projects []*Project) error {
	padding := []int{45, 16, 10, 12, 10, 0}
	head := []string{"Name", "Authors", "Languages", "Tags", "Date", "Source"}
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
	if p.ID != "" {
		name = formatLink(p.Name, p.ID)
	}
	date := p.Time().UTC().Format("2006-01-02")
	links := make([]string, 0, len(p.Source))
	for _, s := range p.Source {
		label, err := GetURLLabel(s)
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
	"github.com":                   "GitHub",
	"gist.github.com":              "GitHub Gist",
	"gitlab.com":                   "GitLab",
	"sourceforge.net":              "SourceForge",
	"git.code.sf.net":              "SourceForge",
	"bitbucket.org":                "Bitbucket",
	"code.google.com":              "Google Code",
	"googlecode.com":               "Google Code",
	"sr.ht":                        "sourcehut",
	"softwareheritage.org":         "Software Heritage",
	"archive.softwareheritage.org": "Software Heritage archive",
	"metacpan.org":                 "CPAN",
	"crates.io":                    "crates.io",
	"hub.docker.com":               "Docker Hub",
	"package.elm-lang.org":         "Elm Packages",
	"hackage.haskell.org":          "Hackage",
	"hex.pm":                       "Hex",
	"mvnrepository.com":            "Maven",
	"npmjs.com":                    "npm",
	"nuget.org":                    "NuGet",
	"opam.ocaml.org":               "opam",
	"packagist.org":                "Packagist",
	"pkg.go.dev":                   "pkg.go.dev",
	"pypi.org":                     "PyPI",
	"rubygems.org":                 "RubyGems",
	"en.wikipedia.org":             "Wikipedia",
	"esolangs.org":                 "Esolang",
	"progopedia.com":               "Progopedia",
	"stackoverflow.com":            "Stack Overflow",
	"codegolf.stackexchange.com":   "Code Golf",
	"codewars.com":                 "Codewars",
	"projecteuler.net":             "Project Euler",
	"rosettacode.org":              "Rosetta Code",
	"adventofcode.com":             "Advent of Code",
	"yukicoder.me":                 "yukicoder",
	"spoj.com":                     "Sphere Online Judge",
	"acmicpc.net":                  "Baekjoon Online Judge",
	"help.acmicpc.net":             "Baekjoon Online Judge",
	"code.activestate.com":         "ActiveState Code",
	"pastebin.com":                 "Pastebin",
	"whitespace.pastebin.com":      "Pastebin",
	"ideone.com":                   "Ideone",
	"drive.google.com":             "Google Drive",
	"compsoc.dur.ac.uk":            "Durham CompSoc",
	"news.ycombinator.com":         "Hacker News",
	"slashdot.org":                 "Slashdot",
	"twitter.com":                  "Twitter",
	"what.thedailywtf.com":         "What the Daily WTF?",
}

func GetURLLabel(rawURL string) (string, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}
	host := u.Hostname()
	if host == "web.archive.org" && strings.HasPrefix(u.Path, "/web/") {
		path := strings.TrimPrefix(u.Path, "/web/")
		if i := strings.IndexByte(path, '/'); i != -1 {
			label, err := GetURLLabel(path[i+1:])
			if err != nil {
				return "", err
			}
			return label + " (archive)", nil
		}
	}
	query := u.Query()
	if host == "archive.softwareheritage.org" && query.Has("origin_url") {
		label, err := GetURLLabel(query.Get("origin_url"))
		if err != nil {
			return "", err
		}
		return label + " (Software Heritage archive)", nil
	}
	if strings.HasSuffix(host, ".googlecode.com") {
		return "Google Code", nil
	}
	if host == "code.google.com" && strings.HasPrefix(u.Path, "/archive/") {
		return "Google Code Archive", nil
	}
	if host == "compsoc.dur.ac.uk" && strings.HasPrefix(u.Path, "/archives/whitespace/") {
		return "Mailing list", nil
	}
	if site, ok := pathTrimPrefix("sites.google.com", "/site/", host, u.Path); ok {
		return site + " Google Site", nil
	}
	if host == "docs.google.com" && strings.HasPrefix(u.Path, "/presentation/") {
		return "Google Slides", nil
	}
	if i := strings.IndexByte(host, '.'); i != -1 {
		switch host[:i] {
		case "www", "www2":
			host = host[i+1:]
		}
	}
	if subreddit, ok := pathTrimPrefix("reddit.com", "/r/", host, u.Path); ok {
		return "r/" + subreddit, nil
	}
	if label, ok := domainLabels[host]; ok {
		return label, nil
	}
	return host, nil
}

func pathTrimPrefix(wantedHost, pathPrefix, host, path string) (string, bool) {
	if host != wantedHost || !strings.HasPrefix(path, pathPrefix) {
		return "", false
	}
	path = strings.TrimPrefix(path, pathPrefix)
	if i := strings.IndexByte(path, '/'); i != -1 {
		path = path[:i]
	}
	return path, true
}
