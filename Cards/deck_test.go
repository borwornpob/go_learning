package main

import (
	"testing"
	"os"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Error("Expected deck length of 48,  but got", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Error("Expected first card of Ace of Spades, but got", d[0])
	}

	if d[len(d) - 1] != "Kings of Clubs" {
		t.Error("Expected last card of Kings of Clubs, but got", d[len(d) - 1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Error("Expected 52 cards in deck, got ", len(loadedDeck))
	}
	
	os.Remove("_decktesting")

}