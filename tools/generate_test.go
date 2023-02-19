package tools

import "testing"

func TestGetURLLabel(t *testing.T) {
	for i, tt := range []struct {
		URL   string
		Label string
	}{
		{"https://github.com/thaliaarchi/nebula", "GitHub"},
		{"https://web.archive.org/web/20200906224652/https://github.com/marcellippmann/Whitepp/", "GitHub (archive)"},
		{"https://gist.github.com/KeenS/6081b0c802a4e575ddbacb1930680870", "GitHub Gist"},
		{"https://gitlab.com/ft/spaceman", "GitLab"},
		{"https://bitbucket.org/StephenPatrick/whitespace-interpreter", "Bitbucket"},
		{"https://archive.softwareheritage.org/browse/origin/directory/?origin_url=https://bitbucket.org/lifthrasiir/esotope-ws", "Bitbucket (Software Heritage archive)"},
		{"https://sourceforge.net/projects/esco/", "SourceForge"},
		{"https://git.code.sf.net/p/esco/code", "SourceForge"},
		{"http://esco.sourceforge.net/", "esco.sourceforge.net"},
		{"https://web.archive.org/web/20160114185408/https://code.google.com/p/grass-mud-horse/", "Google Code (archive)"},
		{"https://web.archive.org/web/20210624084916/https://code.google.com/archive/p/grass-mud-horse/", "Google Code Archive (archive)"},
		{"https://archive.softwareheritage.org/browse/origin/directory/?origin_url=http://grass-mud-horse.googlecode.com/svn/", "Google Code (Software Heritage archive)"},
		{"https://hackage.haskell.org/package/whitespace-0.4", "Hackage"},
		{"https://vii5ard.github.io/whitespace/", "vii5ard.github.io"},
		{"https://web.archive.org/web/20060517072734/http://compsoc.dur.ac.uk:80/archives/whitespace/2004-May/000047.html", "Mailing list (archive)"},
		{"https://web.archive.org/web/20110212015726/http://hapyli.webs.com:80/", "hapyli.webs.com (archive)"},
		{"https://slashdot.org/story/03/04/01/0332202/new-whitespace-only-programming-language", "Slashdot"},
		{"https://www.reddit.com/r/programming/comments/9nw1e/most_unreadable_programming_language_ever/c0dkzzw/", "r/programming"},
		{"http://pelangchallenge.blogspot.com/2013/09/problem-36-done-in-whitespace.html", "pelangchallenge.blogspot.com"},
		{"https://what.thedailywtf.com/topic/5980/stupid-coding-tricks-sudoku-solver-in-whitespace", "What the Daily WTF?"},
		{"https://web.archive.org/web/20100807004910/http://whitespace.pastebin.com/f761fc4b5", "Pastebin (archive)"},
		{"https://web.archive.org/web/20130510111931/https://sites.google.com/site/res0001/whitespace/programs", "res0001 Google Site (archive)"},
		{"https://docs.google.com/presentation/d/1BeNJ_E_KOLjjdM4Bd3u_96tiIGiRdG9J9orhglcxziw/edit", "Google Slides"},
	} {
		label, err := GetURLLabel(tt.URL)
		if err != nil {
			t.Errorf("#%d: GetURLLabel(%q) err: %v", i, tt.URL, err)
			continue
		}
		if label != tt.Label {
			t.Errorf("#%d: GetURLLabel(%q) = %q, want %q", i, tt.URL, label, tt.Label)
		}
	}
}

var urlAndNameTests = []struct {
	URL                      string
	GitURL, Branch, RepoName string
	RepoNameExt              string
}{
	{"https://github.com/wspace/corpus", "https://github.com/wspace/corpus", "", "corpus", "corpus"},
	{"https://github.com/matsubara0507/whitespace_has/tree/develop", "https://github.com/matsubara0507/whitespace_has", "develop", "whitespace_has", "whitespace_has"},
	{"https://gist.github.com/KeenS/6081b0c802a4e575ddbacb1930680870", "https://gist.github.com/KeenS/6081b0c802a4e575ddbacb1930680870", "", "", ""},
	{"https://gitlab.com/ft/spaceman", "https://gitlab.com/ft/spaceman.git", "", "spaceman", "spaceman"},
	{"https://bitbucket.org/StephenPatrick/whitespace-interpreter", "https://bitbucket.org/StephenPatrick/whitespace-interpreter", "", "whitespace-interpreter", "whitespace-interpreter"},
	{"https://git.code.sf.net/p/esco/code", "https://git.code.sf.net/p/esco/code", "", "esco", "esco"},
	{"https://sourceforge.net/projects/esco/", "", "", "", ""},
	{"https://web.archive.org/web/20200906224652/https://github.com/marcellippmann/Whitepp/", "", "", "", "Whitepp"},
	{"https://archive.softwareheritage.org/browse/origin/directory/?origin_url=https://bitbucket.org/lifthrasiir/esotope-ws", "", "", "", "esotope-ws"},
}

func TestGetGitURL(t *testing.T) {
	for i, tt := range urlAndNameTests {
		gitURL, branch, repoName := GetGitURL(tt.URL)
		if gitURL != tt.GitURL || branch != tt.Branch || repoName != tt.RepoName {
			t.Errorf("#%d: GetGitURL(%q) = %q, %q, %q; want %q, %q, %q", i, tt.URL, gitURL, branch, repoName, tt.GitURL, tt.Branch, tt.RepoName)
		}
	}
}

func TestGetRepoName(t *testing.T) {
	for i, tt := range urlAndNameTests {
		repoName := GetRepoName(tt.URL)
		if repoName != tt.RepoNameExt {
			t.Errorf("#%d: GetRepoName(%q) = %q, want %q", i, tt.URL, repoName, tt.RepoNameExt)
		}
	}
}
