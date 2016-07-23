package model

import "testing"

func TestAceCard(t *testing.T) {

	
	card := newCard("Ace")

	if card.cardValue[0] != 1 || card.cardValue[1] !=10 {
		t.Errorf("Should have contained a card with value %d", 1)
	}

}