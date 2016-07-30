package models

type Hand struct {
	cards []*Card
}

func newHand() *Hand {
	return &Hand{
		cards: []*Card{},
	}
}


func (h *Hand) addCard(c *Card) {
	h.cards = append(h.cards, c)
}