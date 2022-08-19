package Web

import (
	"testing"
	"time"
)

func TestFetchWebPageSimple(t *testing.T) {
	node := FetchWebPage("https://www.google.co.uk")
	

	foundTarget := false
	for _, attr := range node.FirstChild.NextSibling.Attr {
		t.Logf("found attr: %s, %s", attr.Key, attr.Val)
		if attr.Key == "itemtype" && attr.Val == "http://schema.org/WebPage" {
			foundTarget = true
			break
		}
	}

	if !foundTarget {
		t.Errorf("Something wrong with html response")
	}
}

func TestFetchWebPageRateLimit(t *testing.T) {
	start := time.Now().UnixNano()
	for i:=0; i<2; i++ {
		FetchWebPage("https://www.google.co.uk")
	}
	end := time.Now().UnixNano()

	if end - start < 2000000000 {
		t.Errorf("responded too fast")
	}
}