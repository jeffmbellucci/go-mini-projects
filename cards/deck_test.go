package main

import (
	"fmt"
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, but got %v", d[0])
	}
	if d[len(d) - 1] != "King of Clubs" {
		t.Errorf("Expected first card to be King of Clubs, but got %v", d[len(d) - 1])
	}
}

func TestSaveDeckToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")
	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(loadedDeck))
	}
	os.Remove("_decktesting")
}

func TestShuffle(t *testing.T) {
	deck := newDeck()
	fmt.Printf("deck[0]: %v\n", deck[0])
	shuffled := newDeck()
	shuffled.shuffle()
	fmt.Printf("shuffled[0]: %v\n", shuffled[0])
	if deck[0] == shuffled[0] {
		t.Errorf("Expected deck to be shuffled")
	}
}