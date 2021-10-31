package tools

import "testing"

func TestGetURLLabel(t *testing.T) {
	tests := []struct {
		URL   string
		Label string
	}{
		{"https://github.com/andrewarchi/nebula", "GitHub"},
		{"https://gist.github.com/KeenS/6081b0c802a4e575ddbacb1930680870", "GitHub Gist"},
		{"https://gitlab.com/ft/spaceman", "GitLab"},
		{"https://bitbucket.org/StephenPatrick/whitespace-interpreter", "Bitbucket"},
		{"https://sourceforge.net/projects/esco/", "SourceForge"},
		{"https://git.code.sf.net/p/esco/code", "SourceForge"},
		{"http://esco.sourceforge.net/", "esco.sourceforge.net"},
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
	}
	for i, tt := range tests {
		label, err := GetURLLabel(tt.URL)
		if err != nil {
			t.Errorf("#%d: getURLLabel(%q) err: %v", i, tt.URL, err)
			continue
		}
		if label != tt.Label {
			t.Errorf("#%d: getURLLabel(%q) = %q, want %q", i, tt.URL, label, tt.Label)
		}
	}
}
