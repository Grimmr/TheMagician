package YGO

import (
	"os"
	"testing"

	"golang.org/x/net/html"
)

func TestGetTableUrlFromHtml(t *testing.T) {
	file, _ := os.Open("../TestHTML/ygo/Set.html")

	rootNode, err := html.Parse(file)
	if err != nil {
		t.Fatalf("got error when reading HTML: %s", err)
	}

	url := (&YgoBackend{}).getTableUrlFromHtml(rootNode)
	expected := "https://yugipedia.com/wiki/Set_Card_Lists:Spell_Ruler_(TCG-EU)"
	if url != expected {
		t.Errorf("expected '%s', got '%s'", expected, url)
	}
}

func TestGetTableUrlFromHtmlNoEU(t *testing.T) {
	file, _ := os.Open("../TestHTML/ygo/Set2.html")

	rootNode, err := html.Parse(file)
	if err != nil {
		t.Fatalf("got error when reading HTML: %s", err)
	}

	url := (&YgoBackend{}).getTableUrlFromHtml(rootNode)
	expected := "https://yugipedia.com/wiki/Set_Card_Lists:Yu-Gi-Oh!_The_Eternal_Duelist_Soul_promotional_cards_(TCG-NA)"
	if url != expected {
		t.Errorf("expected '%s', got '%s'", expected, url)
	}
}
