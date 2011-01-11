package urlshortener

import "testing"

func TestShorten(t *testing.T) {
	shortenUrl, err := ShortenURL("http://golang.org")
	if err != nil {
		t.Error(err)
	}
	t.Log(shortenUrl)
}

func TestExpand(t *testing.T) {
	expandedUrl, err := ExpandURL("http://goo.gl/9YVb")
	if err != nil {
		t.Error(err)
	}
	t.Log(expandedUrl)
}
