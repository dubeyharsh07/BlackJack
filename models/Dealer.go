package models

type Dealer struct {
	deck *Deck
	hand *Hand
}

func setupNewDealer(deck *Deck) *Dealer {
	return &Dealer{deck: deck, hand: newHand()}
}

func (d *Dealer) dealCard() *Card {
	return d.deck.nextCard()
}

func (d *Dealer) selfDealOfCard() {
	d.hand.addCard(d.dealCard())
}
