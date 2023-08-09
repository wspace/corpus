package tools

import (
	"bytes"
	"encoding/json"
	"fmt"
	urlpkg "net/url"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/thaliaarchi/browser/jsonutil"
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
	Submodules        []struct {
		Path   string `json:"path"`
		URL    string `json:"url"`
		Branch string `json:"branch,omitempty"`
	} `json:"submodules,omitempty"`
	Packages []struct {
		Name     string `json:"name"`
		Registry string `json:"registry"`
		URL      string `json:"url"`
	} `json:"packages,omitempty"`
	Relations []struct {
		ID      string `json:"id"`
		Type    string `json:"type"`
		Commit  string `json:"commit,omitempty"`
		Release string `json:"release,omitempty"`
	} `json:"relations,omitempty"`
	Challenges []string `json:"challenges,omitempty"`
	Bounds     *struct {
		Precision      string `json:"precision,omitempty"`
		ArgPrecision   string `json:"arg_precision,omitempty"`   // usually the same as precision
		LabelPrecision string `json:"label_precision,omitempty"` // usually the same as precision
		StackCap       any    `json:"stack_cap,omitempty"`       // int or string
		CallStackCap   any    `json:"call_stack_cap,omitempty"`  // int or string
		HeapMin        any    `json:"heap_min,omitempty"`        // int or string
		HeapMax        any    `json:"heap_max,omitempty"`        // int or string
		HeapCap        any    `json:"heap_cap,omitempty"`        // int or string
		LabelCap       any    `json:"label_cap,omitempty"`       // int or string
		InstructionCap any    `json:"instruction_cap,omitempty"` // int or string
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
			Push      any `json:"push,omitempty"`      // string or []string
			Dup       any `json:"dup,omitempty"`       // string or []string
			Copy      any `json:"copy,omitempty"`      // string or []string
			Swap      any `json:"swap,omitempty"`      // string or []string
			Drop      any `json:"drop,omitempty"`      // string or []string
			Slide     any `json:"slide,omitempty"`     // string or []string
			Add       any `json:"add,omitempty"`       // string or []string
			Sub       any `json:"sub,omitempty"`       // string or []string
			Mul       any `json:"mul,omitempty"`       // string or []string
			Div       any `json:"div,omitempty"`       // string or []string
			Mod       any `json:"mod,omitempty"`       // string or []string
			Store     any `json:"store,omitempty"`     // string or []string
			Retrieve  any `json:"retrieve,omitempty"`  // string or []string
			Label     any `json:"label,omitempty"`     // string or []string
			Call      any `json:"call,omitempty"`      // string or []string
			Jmp       any `json:"jmp,omitempty"`       // string or []string
			Jz        any `json:"jz,omitempty"`        // string or []string
			Jn        any `json:"jn,omitempty"`        // string or []string
			Ret       any `json:"ret,omitempty"`       // string or []string
			End       any `json:"end,omitempty"`       // string or []string
			Printc    any `json:"printc,omitempty"`    // string or []string
			Printi    any `json:"printi,omitempty"`    // string or []string
			Readc     any `json:"readc,omitempty"`     // string or []string
			Readi     any `json:"readi,omitempty"`     // string or []string
			Shuffle   any `json:"shuffle,omitempty"`   // string or []string
			DumpStack any `json:"dumpstack,omitempty"` // string or []string
			DumpHeap  any `json:"dumpheap,omitempty"`  // string or []string
			DumpTrace any `json:"dumptrace,omitempty"` // string or []string
		} `json:"mnemonics,omitempty"`
		Macros []struct {
			Name    string    `json:"name"`
			Args    []string  `json:"args,omitempty"`
			Replace *[]string `json:"replace,omitempty"`
			Notes   string    `json:"notes,omitempty"`
		} `json:"macros,omitempty"`
		Patterns               map[string]string `json:"patterns,omitempty"`
		CaseSensitiveMnemonics *bool             `json:"case_sensitive_mnemonics,omitempty"`
		InstructionDelimiter   string            `json:"instruction_delimiter,omitempty"`
		InstructionWrap        *bool             `json:"instruction_wrap,omitempty"`
		LineComments           []string          `json:"line_comments,omitempty"`
		BlockComments          []struct {
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
	BuildErrors string `json:"build_errors,omitempty"`
	RunErrors   string `json:"run_errors,omitempty"`
	Scripts     *struct {
		Run         *Script `json:"run,omitempty"`
		Compile     *Script `json:"compile,omitempty"`
		Assemble    *Script `json:"assemble,omitempty"`
		Disassemble *Script `json:"disassemble,omitempty"`
	} `json:"scripts,omitempty"`
	Commands []Command `json:"commands,omitempty"`
	Notes    string    `json:"notes,omitempty"`
	TODO     string    `json:"todo,omitempty"`
}

type Script struct {
	Bin   ScriptBin `json:"bin"`
	Args  []string  `json:"args,omitempty"`
	Notes string    `json:"notes,omitempty"`
}

type ScriptBin struct {
	// Path variant
	Path string `json:"path,omitempty"`
	// Jar variant
	Jar  string `json:"jar,omitempty"`
	Main string `json:"main,omitempty"`
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
	Short         any    `json:"short,omitempty"` // -s
	Long          string `json:"long,omitempty"`  // --long
	Bare          string `json:"bare,omitempty"`  // bare
	Required      *bool  `json:"required,omitempty"`
	RepeatAllowed *bool  `json:"repeat_allowed,omitempty"`
	Arg           string `json:"arg,omitempty"`
	ArgRequired   *bool  `json:"arg_required,omitempty"`
	Type          string `json:"type,omitempty"`
	Default       any    `json:"default,omitempty"`
	Min           int    `json:"min,omitempty"`
	Desc          string `json:"desc,omitempty"`
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
		"2006-01-02 15:04 -0700",
		"2006-01-02 15:04",
		"2006-01-02",
		"2006",
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
	date := p.Date
	if len(date) != 4 {
		date = p.Time().UTC().Format("2006-01-02")
	}
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
	"github.com":                             "GitHub",
	"gist.github.com":                        "GitHub Gist",
	"gitlab.com":                             "GitLab",
	"sourceforge.net":                        "SourceForge",
	"git.code.sf.net":                        "SourceForge",
	"bitbucket.org":                          "Bitbucket",
	"code.google.com":                        "Google Code",
	"googlecode.com":                         "Google Code",
	"sr.ht":                                  "sourcehut",
	"codeberg.org":                           "Codeberg",
	"softwareheritage.org":                   "Software Heritage",
	"archive.softwareheritage.org":           "Software Heritage archive",
	"bitbucket-archive.softwareheritage.org": "Software Heritage Mercurial Bitbucket archive",
	"metacpan.org":                           "CPAN",
	"crates.io":                              "crates.io",
	"hub.docker.com":                         "Docker Hub",
	"package.elm-lang.org":                   "Elm Packages",
	"hackage.haskell.org":                    "Hackage",
	"hex.pm":                                 "Hex",
	"mvnrepository.com":                      "Maven",
	"npmjs.com":                              "npm",
	"nuget.org":                              "NuGet",
	"opam.ocaml.org":                         "opam",
	"packagist.org":                          "Packagist",
	"pkg.go.dev":                             "pkg.go.dev",
	"pypi.org":                               "PyPI",
	"rubygems.org":                           "RubyGems",
	"en.wikipedia.org":                       "Wikipedia",
	"esolangs.org":                           "Esolang",
	"progopedia.com":                         "Progopedia",
	"stackoverflow.com":                      "Stack Overflow",
	"codegolf.stackexchange.com":             "Code Golf",
	"codewars.com":                           "Codewars",
	"projecteuler.net":                       "Project Euler",
	"rosettacode.org":                        "Rosetta Code",
	"adventofcode.com":                       "Advent of Code",
	"yukicoder.me":                           "yukicoder",
	"spoj.com":                               "Sphere Online Judge",
	"acmicpc.net":                            "Baekjoon Online Judge",
	"help.acmicpc.net":                       "Baekjoon Online Judge",
	"code.activestate.com":                   "ActiveState Code",
	"pastebin.com":                           "Pastebin",
	"whitespace.pastebin.com":                "Pastebin",
	"ideone.com":                             "Ideone",
	"sites.google.com":                       "Google Sites",
	"drive.google.com":                       "Google Drive",
	"amazon.com":                             "Amazon",
	"speakerdeck.com":                        "Speaker Deck",
	"compsoc.dur.ac.uk":                      "Durham CompSoc",
	"news.ycombinator.com":                   "Hacker News",
	"slashdot.org":                           "Slashdot",
	"twitter.com":                            "Twitter",
	"what.thedailywtf.com":                   "What the Daily WTF?",
}

var subdomainLabels = map[string]string{
	"github.io":       "GitHub Pages",
	"gitlab.io":       "GitLab Pages",
	"sourceforge.net": "SourceForge site",
	"codeberg.page":   "Codeberg Pages",
	"googlecode.com":  "Google Code",
	"readthedocs.io":  "Read the Docs",
	"blogspot.com":    "Blogger",
}

func GetURLLabel(url string) (string, error) {
	u, err := urlpkg.Parse(url)
	if err != nil {
		return "", err
	}
	if archivedURL, archiveLabel := unwrapArchive(url, u); archivedURL != "" {
		label, err := GetURLLabel(archivedURL)
		if err != nil {
			return "", err
		}
		return label + " (" + archiveLabel + ")", nil
	}
	host := u.Hostname()
	if host == "code.google.com" && strings.HasPrefix(u.Path, "/archive/") {
		return "Google Code", nil
	}
	if host == "docs.google.com" && strings.HasPrefix(u.Path, "/presentation/") {
		return "Google Slides", nil
	}
	if host == "compsoc.dur.ac.uk" && strings.HasPrefix(u.Path, "/archives/whitespace/") {
		return "Mailing list", nil
	}
	if host == "github.com" && strings.HasPrefix(u.Path, "/thaliaarchi/repo-archival/") {
		return "repo-archival", nil
	}
	dots := strings.Count(host, ".")
	if dots == 2 {
		if label, ok := subdomainLabels[host[strings.IndexByte(host, '.')+1:]]; ok {
			return label, nil
		}
	}
	if dots >= 2 {
		i := strings.IndexByte(host, '.')
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

var bitbucketHgArchivePattern = regexp.MustCompile(`^https://bitbucket-archive\.softwareheritage\.org/projects/../([^/]+/[^/]+)\.html$`)

func unwrapArchive(url string, u *urlpkg.URL) (string, string) {
	host := u.Hostname()
	if host == "web.archive.org" && strings.HasPrefix(u.Path, "/web/") {
		path := strings.TrimPrefix(u.Path, "/web/")
		if i := strings.IndexByte(path, '/'); i != -1 {
			return path[i+1:], "archive"
		}
	}
	query := u.Query()
	if host == "archive.softwareheritage.org" && query.Has("origin_url") {
		return query.Get("origin_url"), "git archive"
	}
	if match := bitbucketHgArchivePattern.FindStringSubmatch(url); match != nil {
		return "https://bitbucket.org/" + match[1], "hg archive"
	}
	return "", ""
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

var (
	githubPattern      = regexp.MustCompile(`^(https://github\.com/[^/]+/([^/]+))(?:/tree/([^/]+))?$`)
	githubGistPattern  = regexp.MustCompile(`^https://gist\.github\.com/[^/]+/[^/]+$`)
	gitlabPattern      = regexp.MustCompile(`^(https://gitlab\.com/[^/]+/([^/]+))(?:/-/tree/([^/]+))?$`)
	bitbucketPattern   = regexp.MustCompile(`^(https://bitbucket\.org/[^/]+/([^/]+))(?:/src/([^/]+))?$`)
	sourceForgePattern = regexp.MustCompile(`^https://git\.code\.sf\.net/p/([^/]+)/code$`)
	codebergPattern    = regexp.MustCompile(`^https://codeberg\.org/[^/]+/([^/]+)$`)
)

func GetGitURL(url string) (gitURL, branch, repoName string) {
	if match := githubPattern.FindStringSubmatch(url); match != nil {
		return match[1], match[3], match[2]
	}
	if githubGistPattern.MatchString(url) {
		return url, "", ""
	}
	if match := gitlabPattern.FindStringSubmatch(url); match != nil {
		return match[1] + ".git", match[3], match[2]
	}
	if match := bitbucketPattern.FindStringSubmatch(url); match != nil {
		return match[1], match[3], match[2]
	}
	if match := sourceForgePattern.FindStringSubmatch(url); match != nil {
		return url, "", match[1]
	}
	if match := codebergPattern.FindStringSubmatch(url); match != nil {
		return url, "", match[1]
	}
	return "", "", ""
}

func (p *Project) RepoName() string {
	if len(p.Source) == 0 {
		return ""
	}
	url := p.Source[0]
	if strings.HasPrefix(p.Source[0], "https://github.com/wspace/") {
		if len(p.Source) < 2 {
			return ""
		}
		url = p.Source[1]
	}
	return GetRepoName(url)
}

func GetRepoName(url string) string {
	u, err := urlpkg.Parse(url)
	if err != nil {
		return ""
	}
	if archivedURL, _ := unwrapArchive(url, u); archivedURL != "" {
		url = strings.TrimSuffix(archivedURL, "/")
	}
	_, _, repoName := GetGitURL(url)
	return repoName
}
