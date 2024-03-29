package main

import (
	"os"
	"testing"
)

// Test function names should begin with Test (note the capital T)
func TestNewDeck(t *testing.T) {
	// The variable t represents a test handler that you'll
	// use to indicate what went wrong during testing.
	d := newDeck()

	// Instead of using assert statements, we're just using plain ol' if statements here...
	//  - Note that the Udemy tutorial probably didn't have an assertion
	//	  library available at the time?
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16 but got %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades but got %s", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs but got %s", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck length of 16 but got %d", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
