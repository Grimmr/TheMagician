package YGO

import (
	"testing"
	"golang.org/x/net/html"
	"os"
)

func TestFindInfoColumn(t *testing.T) {
	file, _ := os.Open("../TestHTML/ygo/Card.html")
	
	rootNode, err := html.Parse(file)
	if err != nil {
		t.Fatalf("got error when reading HTML: %s", err)
	}
	
	found := findInfoColumn(rootNode)
	if found == nil {
		t.Fatalf("couldn't find target node")
	}
}