package main

func main() {

	newCards := readDeckFromFile("deckFile")
	d1 := newCards
	newCards.shuffle()
	d1.print()
	print("----------")
	newCards.print()

}
