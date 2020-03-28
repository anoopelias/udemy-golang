package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Mismatched deck size, expected 16, got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("First card should be %v", "Ace of Spades")
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Last card should be %v", "Four of Clubs")
	}

}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	df := newDeckFromFile("_decktesting")

	if len(df) != 16 {
		t.Errorf("Expecter length to be 16, got %v", len(df))
	}

}
