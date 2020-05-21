package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

//Create a new type deck
//which is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Diamonds", "Hearts", "Spades", "Clovers"}
	cardValues := []string{"Ace", "two", "three", "four"} //, "five", "six", "seven", "eight", "nine", "ten", "jack", "king", "Queen"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func readDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("error", err)
		return nil
	}
	return (deck(strings.Split(string(bs), ",")))
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) toString() string {
	return (strings.Join([]string(d), ","))
}

func (d deck) saveToFile(filename string) error {
	return (ioutil.WriteFile(filename, []byte(d.toString()), 0644))
}

func (d deck) shuffle() {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPos := r.Intn(len(d) - 1)
		d[i], d[newPos] = d[newPos], d[i]
	}

}
