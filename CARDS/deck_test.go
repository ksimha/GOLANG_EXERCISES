package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16 but got %d", len(d))
	}

	if d[0] != "Ace of Diamonds" {
		t.Errorf("Expected card Ace of Diamonds but found %s", d[0])
	}

	if d[len(d)-1] != "four of Clovers" {
		t.Errorf("Expected card four of Clovers but found %s", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_deckTestFile")
	d := newDeck()
	d.saveToFile("_deckTestFile")

	loadedDeck := readDeckFromFile("_deckTestFile")
	if len(loadedDeck) != 16 {
		t.Errorf("Expecting 16 cards in deck, got %d", len(loadedDeck))
	}

	os.Remove("_deckTestFile")

	emptyDeck := readDeckFromFile("_deckTestFile")
	if emptyDeck != nil {
		t.Errorf("Expected nil deck but got value")
	}

}

func TestDeal(t *testing.T) {
	d := newDeck()

	cardsInHand, remainingCardsInDeck := deal(d, 12)

	if len(cardsInHand) != 12 {
		t.Errorf("cards in hand is expected to be 12 but found to be %d", len(cardsInHand))
	}
	if len(remainingCardsInDeck) != 4 {
		t.Errorf("remaining cards in deck should be 4 but is %d", len(remainingCardsInDeck))
	}

}

func TestShuffle(t *testing.T) {
	deck := newDeck()
	backupDeck := newDeck()

	deck.shuffle()

	if (deck[0] == backupDeck[0]) ||
		deck[len(deck)-1] == backupDeck[len(backupDeck)-1] {
		t.Errorf("Not shuffled properly")
	}
}
