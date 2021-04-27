package main

import (
	"testing"
)

func TestGetURLLabel(t *testing.T) {
	tests := []struct {
		URL   string
		Label string
	}{
		{"https://github.com/andrewarchi/nebula", "GitHub"},
		{"https://gitlab.com/ft/spaceman", "GitLab"},
		{"https://web.archive.org/web/20141011193201/http://compsoc.dur.ac.uk/archives/whitespace/2004-May/000047.html", "Mailing list"},
		{"https://web.archive.org/web/20110212015726/http://hapyli.webs.com/", "hapyli.webs.com"},
		{"https://web.archive.org/web/20061209092911/http://www.cs.newcastle.edu.au/~c3018900/pywhitespace.tar.bz2", "Newcastle"},
	}
	for i, tt := range tests {
		label, err := getURLLabel(tt.URL)
		if err != nil {
			t.Errorf("#%d: getURLLabel(%s) err: %v", i, tt.URL, err)
			continue
		}
		if label != tt.Label {
			t.Errorf("#%d: getURLLabel(%s) = %s, want %s", i, tt.URL, label, tt.Label)
		}
	}
}
