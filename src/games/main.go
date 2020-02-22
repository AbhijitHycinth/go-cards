package main

import (
	"fmt"
	"games/poker"
)

func main() {
	deck := poker.NewDeck()
	fmt.Println([]string(deck))
}
