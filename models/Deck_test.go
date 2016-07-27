package model

import "testing"

func TestEmptyDeck(t *testing.T) {

}

func TestDeckInitialization(t *testing.T) {
	
	deck := initializeDeck()

	if len(deck.cards) != 52 {
		t.Errorf("Expected initailized size of deck is 52, got", len(deck))
	}
}