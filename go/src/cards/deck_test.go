package main

import (
	"math/rand"
	"os"
	"reflect"
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

func TestDeal(t *testing.T) {
	cards := newDeck()
	hand, remaining := deal(cards, 5)

	if len(hand) != 5 {
		t.Errorf("expected hand length of 5, got %d", len(hand))
	}

	if len(remaining) != 47 {
		t.Errorf("expected remaining deck length of 47, got %d", len(remaining))
	}

	if hand[0] != "Ace of Spades" {
		t.Errorf("expected first hand card to be Ace of Spades, got %q", hand[0])
	}

	if remaining[0] != "Six of Spades" {
		t.Errorf("expected first remaining card to be Six of Spades, got %q", remaining[0])
	}
}

func TestToString(t *testing.T) {
	cards := deck{"Ace of Spades", "Two of Hearts"}
	got := cards.toString()

	if got != "Ace of Spades,Two of Hearts" {
		t.Errorf("expected joined deck string, got %q", got)
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	filename := "test_deck.txt"
	cards := deck{"Ace of Spades", "Two of Hearts"}

	if err := cards.saveToFile(filename); err != nil {
		t.Errorf("saveToFile returned error: %v", err)
	}
	defer os.Remove(filename)

	loaded := newDeckFromFile(filename)
	if !reflect.DeepEqual(loaded, cards) {
		t.Errorf("expected loaded deck %v, got %v", cards, loaded)
	}
}

func TestShuffleChangesOrder(t *testing.T) {
	cards := newDeck()
	original := append(deck(nil), cards...)

	shuffleWithRand(&cards, rand.New(rand.NewSource(42)))

	if reflect.DeepEqual(cards, original) {
		t.Errorf("expected shuffled deck to differ from the original order")
	}
}
