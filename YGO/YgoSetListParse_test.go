package YGO

import (
	"os"
	"testing"

	"golang.org/x/net/html"
)

func TestGetSetUrlsFromHtml(t *testing.T) {
	file, _ := os.Open("../TestHTML/ygo/SetList.html")

	rootNode, err := html.Parse(file)
	if err != nil {
		t.Fatalf("got error when reading HTML: %s", err)
	}

	found := (&YgoBackend{}).getSetUrlsFromHtml(rootNode)

	expectedRow := map[string]string{"date": "2004-August-13", "name": "Movie Pack", "url": "https://yugipedia.com/wiki/Movie_Pack"}

	for _, row := range found {
		foundAll := true
		for k, v := range expectedRow {
			if row[k] != v {
				foundAll = false
			}
		}
		if foundAll {
			return
		}
	}
	t.Errorf("couldn't find targert row")
}
