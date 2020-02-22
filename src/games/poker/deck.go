package poker

import (
	"fmt"
	"games/poker/cards"
)

//Deck to represent the deck of cards
type Deck []string

// NewDeck Mehod to create a new deck for the game
func NewDeck() Deck {
	var newDeck Deck
	suits := cards.NewSuit()
	cards := cards.NewCards()

	for _, suit := range suits {
		for _, card := range cards {
			newDeck = append(newDeck, cardString(card, suit))
		}
	}

	fmt.Println("Successfully created")
	return newDeck
}

func cardString(card, suit string) string {
	return card + " of " + suit
}
