package models

import "testing"

func TestAceCard(t *testing.T) {

	
	card := newCard("Ace")

	if card.cardValue[0] != 1 || card.cardValue[1] !=10 {
		t.Errorf("Should have contained a card with value %d", 1)
	}

}

func TestCardWithValueAndType(t *testing.T) {
	cards := []struct {
		inputType string
		expectedValue int
	}{
		{"Card2", 2},
		//{Card3, 3},
		{"Card4", 4},
		{"Card5", 5},
		//{Card6, 6},
		{"Card7", 7},
		{"Card8", 8},
		//{Card9, 9},
		{"Card10", 10},
		{"Jack", 10},
		{"Queen",10},
		{"King", 10},
	}

	for _, element := range cards {
		card := newCard(element.inputType)

		if card.cardType != element.inputType {
			t.Errorf("Expected type to be '%s', but got '%s'", element.inputType, card.cardType)
		}

		if card.cardValue[0] != element.expectedValue {
			t.Errorf("Expected value to be %d, but got %d", element.expectedValue, card.cardValue[0])
		}
	}

}