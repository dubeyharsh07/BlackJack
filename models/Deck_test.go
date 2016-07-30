package models

import "testing"



func TestDeckInitialization(t *testing.T) {
	
	deck := initializeDeck()

	if len(deck.cards) != 52 {
		t.Errorf("Expected initailized size of deck is 52, got %d", len(deck))
	}
}

func TestEmptyDeck(t *testing.T) {

	deck := emptyDeck()
	cardCount := len(deck.cards)
	if cardCount != 0 {
		t.Errorf("Expected card count of empty deck to be 0, but got %d", cardCount)
	}
}


func TestShouldRetrunErrorWhenAddingNewCardToInitializedDeck(t *testing.T) {
	deck := initializeDeck()
	err := deck.addCard("King")

	if err == nil {
		t.Errorf("Expected error from addcard method.")
	}
}

func TestShouldRetrunErrorWhenAdding5thCardOfSameTypeIntoADeck(t *testing.T) {
	deck := emptyDeck()
	for i := 1; i <= 4; i++ {
		err := deck.addCard("Jack")
		if err != nil {
			t.Errorf("Expected error from addcard method.")
		}
	}

	err := deck.addCard("Jack")
	if err == nil {
		t.Errorf("Expected error from addcard method.")
	}

	
}

