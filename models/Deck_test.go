package model

import "testing"

func TestDeckInitialization(t *testing.T) {
	
	deck := initializeDeck()

	if len(deck) != 52 {
		t.Errorf("Expected initailized size of deck is 52, got", len(deck))
	}
}