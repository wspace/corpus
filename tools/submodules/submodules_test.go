package main

import "testing"

func TestGetGitURL(t *testing.T) {
	tests := []struct {
		URL, GitURL, Branch string
	}{
		{"https://github.com/wspace/corpus", "https://github.com/wspace/corpus", ""},
		{"https://github.com/matsubara0507/whitespace_has/tree/develop", "https://github.com/matsubara0507/whitespace_has", "develop"},
		{"https://gist.github.com/KeenS/6081b0c802a4e575ddbacb1930680870", "https://gist.github.com/KeenS/6081b0c802a4e575ddbacb1930680870", ""},
		{"https://gitlab.com/ft/spaceman", "https://gitlab.com/ft/spaceman.git", ""},
		{"https://bitbucket.org/StephenPatrick/whitespace-interpreter", "https://bitbucket.org/StephenPatrick/whitespace-interpreter", ""},
		{"https://git.code.sf.net/p/esco/code", "https://git.code.sf.net/p/esco/code", ""},
		{"https://sourceforge.net/projects/esco/", "", ""},
	}
	for i, tt := range tests {
		gitURL, branch := getGitURL(tt.URL)
		if gitURL != tt.GitURL {
			t.Errorf("#%d: getGitURL(%q) = %q, %q; want %q, %q", i, tt.URL, gitURL, branch, tt.GitURL, tt.Branch)
		}
	}
}
