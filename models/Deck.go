package model

type Deck struct {
	cards []Card
}

func initializeDeck() []*Card {

	deck := []Deck{}
	for key,value := range cardValues {
		card := newCard(key)
		deck = append(deck,card)
		deck = append(deck,card)
		deck = append(deck,card)
		deck = append(deck,card)
	}
	return deck.cards
}
