package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if (len(d) != 52) {
		t.Errorf("Expected deck length of 52, got %v", len(d))
	}

	if (d[0] != "Ace of Hearts") {
		t.Errorf("Expected first card 'Ace of Hearts, got %v", d[0])
	}

	if (d[len(d) -1] != "King of Diamonds") {
		t.Errorf("Expected last card 'King of Diamonds', got %v", d[len(d) -1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	d := newDeck()
	d.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")

	if (len(loadedDeck) != 52){
		t.Errorf("Expected deck length of 52, got %v", len(d))
	}
	os.Remove("_decktesting")
}