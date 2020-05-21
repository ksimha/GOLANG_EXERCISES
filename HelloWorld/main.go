package main

import "fmt"

type deck []string

func main() {
	cards := deck{}
	cards = newDeck()

	cards.print()
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}

}
func newDeck() deck {
	return deck{"abc", "def"}
}
