package poker

import (
	"fmt"
	"games/poker/cards"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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

	return newDeck
}

func cardString(card, suit string) string {
	return card + " of " + suit
}

//Print method to print a deck of cards
func (d Deck) Print() {
	for _, card := range d {
		fmt.Println(card)
	}

	fmt.Println("Total number of cards in the deck: ", len(d))
}

//WriteDeckToFile method to write the deck to a local file
func (d Deck) WriteDeckToFile(fileName string) {
	joinedString := strings.Join(d, ",\n")
	err := ioutil.WriteFile(fileName, []byte(joinedString), 0777)

	if err != nil {
		fmt.Println("Error while saving file:", fileName)
		os.Exit(1)
	}
}

//ReadDeckFromFile method to read the deck from a local file
func ReadDeckFromFile(fileName string) Deck {
	byteSlice, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error while reading file:", fileName)
		os.Exit(1)
	}
	joinedString := string(byteSlice)
	deck := Deck(strings.Split(joinedString, "\n"))

	return deck
}

//Shuffle method to shuffle deck of cards
func (d Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano()) ///Seeding the random number generator to generate new value everytime

	for i := range d {
		indexToSwap := rand.Intn(len(d) - 1)
		d[i], d[indexToSwap] = d[indexToSwap], d[i] //Swap cards in a single line
	}
}
