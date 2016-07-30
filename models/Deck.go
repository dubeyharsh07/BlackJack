package models


type Deck struct {
	cards []*Card
}

func initializeDeck() *Deck {

	cards := initializeCards()

	deck := &Deck{
		cards: cards,
	}
	return deck
}

func emptyDeck() *Deck {
	return &Deck{
		cards: []*Card{},
	}
}

func (deck *Deck) addCard(cardType string) error {
	if len(deck.cards) >= 52 {
		return errors.New("Deck is full.")
	}

	cardTypeCount := deck.getSimilarCardCount(cardType)
	if cardTypeCount >= 4 {
		return fmt.Errorf("Number of cards of type %s in deck exceeds.", cardType)
	}

	deck.cards = append(deck.cards, newCard(cardType))
	return nil
}

func (deck *Deck) getSimilarCardCount(cardType string) int {
	count := 0
	for _, card := range deck.cards {
		if card.cardType == cardType {
			count++
		}
	}
	return count
}

func initializeCards() []*Card {
	var cards []*Card
	for cardType, _ := range cardValues {
		cards = append(cards, newCard(cardType))
		cards = append(cards, newCard(cardType))
		cards = append(cards, newCard(cardType))
		cards = append(cards, newCard(cardType))
	}

	return cards
}

func (d *Deck) nextCard() *Card {
	card := d.cards[0]
	d.cards = d.cards[1:]
	return card
}