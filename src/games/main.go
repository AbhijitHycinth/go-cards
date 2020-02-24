package main

import (
	"games/poker"
)

func main() {
	deck := poker.ReadDeckFromFile("new_file.txt")
	deck.Shuffle()
	deck.Print()
}
