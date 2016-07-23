package model

import "fmt" 

	
var cardValues map[string][]int {
	Ace : []{1,10}
	Card2 : []{2}
	Card3 : []{3}
	Card4 : []{4}
	Card5 : []{5}
	Card6 : []{6}
	Card7 : []{7}
	Card8 : []{8}
	Card9 : []{9}
	Card10 : []{10}
	Jack : []{10}
	Queen : []{10}
	King : []{10}
}

type Card struct {
	cardValue []int
	cardType string
}

