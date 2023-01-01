package YGO

import (
	"os"
	"testing"

	"golang.org/x/net/html"
)

func TestGetCardDataFromHtmlCard1(t *testing.T) {
	file, _ := os.Open("../TestHTML/ygo/Card.html")

	rootNode, err := html.Parse(file)
	if err != nil {
		t.Fatalf("got error when reading HTML: %s", err)
	}

	card := getCardDataFromHtml(rootNode)
	expected := map[string]string{"name": "Blue-Eyes Toon Dragon", "password": "53183600", "type": "Monster"}

	expectedLen := len(expected)
	actualLen := len(card)
	if expectedLen != actualLen {
		t.Fatalf("expected length %d, got %d", expectedLen, actualLen)
	}

	for k, v := range expected {
		if card[k] != v {
			t.Errorf("for key %s: expected %s, got %s", k, v, card[k])
		}
	}
}

func TestGetCardDataFromHtmlCard2(t *testing.T) {
	file, _ := os.Open("../TestHTML/ygo/Card2.html")

	rootNode, err := html.Parse(file)
	if err != nil {
		t.Fatalf("got error when reading HTML: %s", err)
	}

	card := getCardDataFromHtml(rootNode)
	expected := map[string]string{"name": "Toon Dark Magician", "password": "21296502", "type": "Monster"}

	expectedLen := len(expected)
	actualLen := len(card)
	if expectedLen != actualLen {
		t.Fatalf("expected length %d, got %d", expectedLen, actualLen)
	}

	for k, v := range expected {
		if card[k] != v {
			t.Errorf("for key %s: expected %s, got %s", k, v, card[k])
		}
	}
}

func TestReg1(t *testing.T) {
	file, _ := os.Open("../TestHTML/ygo/Card3.html")

	rootNode, err := html.Parse(file)
	if err != nil {
		t.Fatalf("got error when reading HTML: %s", err)
	}

	card := getCardDataFromHtml(rootNode)
	expected := map[string]string{"name": "Marauding Captain", "password": "02460565", "type": "Monster"}

	expectedLen := len(expected)
	actualLen := len(card)
	if expectedLen != actualLen {
		t.Fatalf("expected length %d, got %d", expectedLen, actualLen)
	}

	for k, v := range expected {
		if card[k] != v {
			t.Errorf("for key %s: expected %s, got %s", k, v, card[k])
		}
	}
}
