package main 

import (
	"testing"

	m "github.com/dubeyharsh07/BlackJack_Exercise/BlackJack/models"
)

func TestInitializeOfGame(t *testing.T) {

	deck := initializeDeck()

	game := newGame(deck)
	//game.start() returns the output

	// add test condition
}
