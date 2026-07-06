package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	cards := newDeck()

	if len(cards) != 52 {
		t.Errorf("expected deck length of 52, got %d", len(cards))
	}

	if cards[0] != "Ace of Spades" {
		t.Errorf("expected first card to be Ace of Spades, got %q", cards[0])
	}

	if cards[len(cards)-1] != "King of Clubs" {
		t.Errorf("expected last card to be King of Clubs, got %q", cards[len(cards)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("test_deck.txt") // clean up before test

	deck := newDeck()
	deck.saveToFile("test_deck.txt")

	loadedDeck := newDeckFromFile("test_deck.txt")

	if len(loadedDeck) != 52 {
		t.Errorf("expected loaded deck length of 52, got %d", len(loadedDeck))
	}

	os.Remove("test_deck.txt") // clean up after test
}
