package models

var cardValues = map[string][]int{
	"Ace" : []int{1,10},
	"Card2" : []int{2},
	"Card3" : []int{3},
	"Card4" : []int{4},
	"Card5" : []int{5},
	"Card6" : []int{6},
	"Card7" : []int{7},
	"Card8" : []int{8},
	"Card9" : []int{9},
	"Card10" : []int{10},
	"Jack" : []int{10},
	"Queen" : []int{10},
	"King" : []int{10},
}

type Card struct {
	cardValue []int
	cardType string
}

func newCard(t string) *Card {

	newCard := Card{cardType:t, cardValue: cardValues[t]}
	return &newCard
}

