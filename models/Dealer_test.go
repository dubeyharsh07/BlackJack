package model

import "testing"


func TestDealerDealsTwoCards(t *testing.T) {
	dealer := setupNewDealer(initializeDeck())
	card := dealer.dealCard()

	if card == nil {
		t.Errorf("Expected card not to be nil")
	}
}

func TestDealerDealsCardToSelf(t *testing.T) {
	dealer := setupNewDealer(initializeDeck())

	dealer.selfDealOfCard()

	cardsCount := len(dealer.hand.cards) 

	if cardsCount != 1 {
		t.Errorf("Dealer should have 1 card in hand, but got %d", cardsCount)
	}
}
