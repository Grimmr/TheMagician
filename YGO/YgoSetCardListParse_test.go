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

func TestGgtCardUrlsFromTableHtml(t *testing.T) {
	file, _ := os.Open("../TestHTML/ygo/SetCardList.html")

	rootNode, err := html.Parse(file)
	if err != nil {
		t.Fatalf("got error when reading HTML: %s", err)
	}

	cardUrls := (&YgoBackend{}).getCardUrlsFromTableHtml(rootNode)

	if len(cardUrls) != 131 {
		t.Fatalf("expected 131 rows, but got %d", len(cardUrls))
	}

	expectedCard := "https://yugipedia.com/wiki/Blue-Eyes_Toon_Dragon"
	actualCard := cardUrls[0]
	if actualCard != expectedCard {
		t.Errorf("expected %s, got %s", expectedCard, actualCard)
	}

	expectedCard = "https://yugipedia.com/wiki/Ameba"
	actualCard = cardUrls[10]
	if actualCard != expectedCard {
		t.Errorf("expected %s, got %s", expectedCard, actualCard)
	}

	expectedCard = "https://yugipedia.com/wiki/Labyrinth_Wall"
	actualCard = cardUrls[55]
	if actualCard != expectedCard {
		t.Errorf("expected %s, got %s", expectedCard, actualCard)
	}

	expectedCard = "https://yugipedia.com/wiki/Serpent_Night_Dragon"
	actualCard = cardUrls[130]
	if actualCard != expectedCard {
		t.Errorf("expected %s, got %s", expectedCard, actualCard)
	}
}
