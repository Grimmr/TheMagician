package Html

import (
	"os"
	"testing"

	"golang.org/x/net/html"
)

func TestFindNodesWithAttrsSingle(t *testing.T) {
	file, _ := os.Open("../TestHTML/ygo/Card.html")

	rootNode, err := html.Parse(file)
	if err != nil {
		t.Fatalf("got error when reading HTML: %s", err)
	}

	targets := []map[string]string{
		{"title": "Password"}}

	found := FindNodesWithAttrs(rootNode, targets)
	if found == nil {
		t.Fatalf("couldn't find target node")
	}

	if found[0].Attr[0].Val != "/wiki/Password" {
		t.Fatalf("found wrong node")
	}
}

func TestFindNodesWithAttrsMulti(t *testing.T) {
	file, _ := os.Open("../TestHTML/ygo/Card.html")

	rootNode, err := html.Parse(file)
	if err != nil {
		t.Fatalf("got error when reading HTML: %s", err)
	}

	targets := []map[string]string{
		{"class": "mw-body-content", "id": "siteNotice"}}

	found := FindNodesWithAttrs(rootNode, targets)
	if found == nil {
		t.Fatalf("couldn't find target node")
	}

	if found[0].FirstChild.Attr[0].Val != "localNotice" {
		t.Fatalf("found wrong node")
	}
}

func TestFindNodesWithAttrsMultiNode(t *testing.T) {
	file, _ := os.Open("../TestHTML/ygo/Card.html")

	rootNode, err := html.Parse(file)
	if err != nil {
		t.Fatalf("got error when reading HTML: %s", err)
	}

	targets := []map[string]string{
		{"class": "mw-body-content", "id": "siteNotice"},
		{"title": "Password"}}

	found := FindNodesWithAttrs(rootNode, targets)
	if found == nil {
		t.Fatalf("couldn't find target node")
	}

	if found[0].FirstChild.Attr[0].Val != "localNotice" {
		t.Fatalf("found wrong node")
	}

	if found[1].Attr[0].Val != "/wiki/Password" {
		t.Fatalf("found wrong node")
	}
}

func TestFindNodeWithDataSimple(t *testing.T) {
	file, _ := os.Open("../TestHTML/ygo/Dummy.html")

	rootNode, err := html.Parse(file)
	if err != nil {
		t.Fatalf("got error when reading HTML: %s", err)
	}

	targets := []string{"someData"}

	found := FindNodesWithData(rootNode, targets)
	if len(found) != 1 {
		t.Fatalf("expected 1 result, got %d", len(found))
	}

	if found[0].Data != "someData" {
		t.Errorf("found wrong node")
	}
}
