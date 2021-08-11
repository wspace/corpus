package tools

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"
)

type Project struct {
	Name        string   `json:"name"`
	Description string   `json:"description,omitempty"`
	Path        string   `json:"path,omitempty"`
	Authors     []string `json:"authors"`
	License     string   `json:"license,omitempty"`
	Languages   []string `json:"languages"`
	Tags        []string `json:"tags"`
	Date        string   `json:"date"`
	SpecVersion string   `json:"spec_version"`
	Source      []string `json:"source"`
	Package     *struct {
		Name    string `json:"name"`
		Manager string `json:"manager"`
		URL     string `json:"url"`
	} `json:"package,omitempty"`
	Relations []struct {
		Type string `json:"type"`
		ID   string `json:"id"`
	} `json:"relations,omitempty"`
	Bounds *struct {
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
			Name    string `json:"name"`
			Replace string `json:"replace"`
			Notes   string `json:"notes,omitempty"`
		} `json:"macros,omitempty"`
		Patterns                  map[string]string `json:"patterns,omitempty"`
		CaseSensitiveInstructions *bool             `json:"case_sensitive_instructions,omitempty"`
		LineComments              []string          `json:"line_comments,omitempty"`
		BlockComments             []struct {
			Start  string `json:"start"`
			End    string `json:"end"`
			Nested bool   `json:"nested"`
		} `json:"block_comments,omitempty"`
		Indentation      string   `json:"indentation,omitempty"`
		LabelIndentation string   `json:"label_indentation,omitempty"`
		BinaryNumbers    *bool    `json:"binary_numbers,omitempty"`
		Usage            []string `json:"usage,omitempty"`
		Extension        string   `json:"extension,omitempty"`
	} `json:"assembly,omitempty"`
	Mappings []struct {
		Space         string `json:"space"`
		Tab           string `json:"tab"`
		LF            string `json:"lf"`
		SpacesBetween *bool  `json:"spaces_between,omitempty"`
		LineComment   string `json:"line_comment,omitempty"`
		BeforeSTL     *bool  `json:"before_stl,omitempty"`
		IgnoreCase    *bool  `json:"ignore_case,omitempty"`
		Extension     string `json:"extension,omitempty"`
	} `json:"mappings,omitempty"`
	Programs []struct {
		Path        string   `json:"path"`
		Compiled    string   `json:"compiled,omitempty"`
		Inputs      []string `json:"inputs,omitempty"`
		Outputs     []string `json:"outputs,omitempty"`
		Polyglot    []string `json:"polyglot,omitempty"`
		SpecVersion string   `json:"spec_version,omitempty"`
		Generate    string   `json:"generate,omitempty"`
		Desc        string   `json:"desc,omitempty"`
	} `json:"programs,omitempty"`
	Commands []Command `json:"commands,omitempty"`
	Notes    string    `json:"notes,omitempty"`
}

type Command struct {
	Type                string          `json:"type,omitempty"`
	Bin                 string          `json:"bin,omitempty"`
	Dependencies        []string        `json:"dependencies,omitempty"`
	InstallDependencies []string        `json:"install_dependencies,omitempty"`
	Build               string          `json:"build,omitempty"`
	BuildErrors         string          `json:"build_errors,omitempty"`
	Usage               string          `json:"usage,omitempty"`
	RunErrors           string          `json:"run_errors,omitempty"`
	Input               string          `json:"input,omitempty"`
	Output              string          `json:"output,omitempty"`
	Options             []CommandOption `json:"options,omitempty"`
	OptionParse         string          `json:"option_parse,omitempty"`
	Subcommands         []struct {
		Name    string          `json:"name"`
		Desc    string          `json:"desc,omitempty"`
		Usage   string          `json:"usage,omitempty"`
		Options []CommandOption `json:"options,omitempty"`
		Notes   string          `json:"notes,omitempty"`
	} `json:"subcommands,omitempty"`
	Unimplemented []string `json:"unimplemented,omitempty"`
	Notes         string   `json:"notes,omitempty"`
}

type CommandOption struct {
	Short       interface{} `json:"short,omitempty"` // -s
	Long        string      `json:"long,omitempty"`  // --long
	Bare        string      `json:"bare,omitempty"`  // bare
	Arg         string      `json:"arg,omitempty"`
	ArgRequired *bool       `json:"arg_required,omitempty"`
	Type        string      `json:"type,omitempty"`
	Default     interface{} `json:"default,omitempty"`
	Min         interface{} `json:"min,omitempty"`
	Desc        string      `json:"desc,omitempty"`
	Values      []struct {
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

func RenderProjectTable(b *strings.Builder, projects []Project) error {
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
	} else if t, err := time.Parse("2006-01-02 15:04:05", date); err == nil {
		date = t.Format("2006-01-02")
	} else if t, err := time.Parse("2006-01-02 15:04", date); err == nil {
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
	"gist.github.com":            "GitHub Gist",
	"gitlab.com":                 "GitLab",
	"bitbucket.org":              "Bitbucket",
	"sourceforge.net":            "SourceForge",
	"git.code.sf.net":            "SourceForge",
	"hackage.haskell.org":        "Hackage",
	"en.wikipedia.org":           "Wikipedia",
	"progopedia.com":             "Progopedia",
	"codegolf.stackexchange.com": "Code Golf",
	"rosettacode.org":            "Rosetta Code",
	"codewars.com":               "Codewars",
	"code.activestate.com":       "ActiveState Code",
	"pastebin.com":               "Pastebin",
	"whitespace.pastebin.com":    "Pastebin",
	"compsoc.dur.ac.uk":          "CompSoc",
	"cs.newcastle.edu.au":        "Newcastle",
	"dcode.fr":                   "dCode",
	"news.ycombinator.com":       "HN",
	"slashdot.org":               "Slashdot",
	"what.thedailywtf.com":       "What the Daily WTF?",
}

var subSites = map[string]struct{}{
	// "blogspot.com": {},
}

var pathPatterns = map[string]*regexp.Regexp{
	"reddit.com":       regexp.MustCompile("^/(r/[^/]+)"),
	"sites.google.com": regexp.MustCompile("^/site/([^/]+)"),
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
		if match := pattern.FindStringSubmatch(u.Path); len(match) > 1 {
			return match[1], nil
		}
	}
	if label, ok := domainLabels[host]; ok {
		return label, nil
	}
	return host, nil
}
