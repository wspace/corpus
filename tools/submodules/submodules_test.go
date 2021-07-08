package main

import "testing"

func TestGetGitURL(t *testing.T) {
	tests := []struct {
		URL, GitURL string
	}{
		{"https://github.com/wspace/corpus", "https://github.com/wspace/corpus"},
		{"https://gitlab.com/ft/spaceman", "https://gitlab.com/ft/spaceman.git"},
		{"https://sourceforge.net/projects/esco/", "https://git.code.sf.net/p/esco/code"},
	}
	for i, tt := range tests {
		gitURL := getGitURL(tt.URL)
		if gitURL != tt.GitURL {
			t.Errorf("#%d: getGitURL(%q) = %q, want %q", i, tt.URL, gitURL, tt.GitURL)
		}
	}
}
