package poker

import (
	"os"
	"testing"
)

func TestShouldCreateNewDeck(t *testing.T) {
	deck := NewDeck()

	if len(deck) != 52 {
		t.Errorf("\nExpected deck length of 52, got %v", len(deck))
	}

	if deck[0] != "Ace of Heart" {
		t.Errorf("\nExpected to get \"Ace of Heart\", got %v", deck[0])
	}

	if deck[51] != "King of Club" {
		t.Errorf("\nExpected to get \"King of Club\", got %v", deck[51])
	}
}

func TestShouldSaveDeckToFile(t *testing.T) {
	deck := NewDeck()
	fileName := "_testingFile"

	deck.WriteDeckToFile(fileName)
	loadedDeck := ReadDeckFromFile(fileName)

	if len(deck) != len(loadedDeck) {
		t.Errorf("Length mismatch of deck")
	}

	os.Remove(fileName)
}
