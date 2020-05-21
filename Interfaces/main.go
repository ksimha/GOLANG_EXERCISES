package main

import "fmt"

type bot interface {
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}

type test struct {
	testBot bot
}

func (e englishBot) getGreeting() string {
	//very custom logic to generate a Englis greeting ;)
	return "Hello"
}

func (s spanishBot) getGreeting() string {
	//very custom logic to generate a spanish greeting ;)
	return "Hola"
}

/*func (d slicetype) getGreeting() string {
	return "Namaskara"
}*/

func (t test) getGreeting() string {
	return "test"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

type shape interface {
	getArea() float64
}

type triangle struct {
	base   float64
	height float64
}

type square struct {
	sideLen float64
}

func (t triangle) getArea() float64 {
	return 0.5 * t.height * t.base
}

func (s square) getArea() float64 {
	return s.sideLen * s.sideLen
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func main() {

	sq := square{sideLen: 12}
	t := triangle{base: 12, height: 14}

	printArea(sq)
	printArea(t)
}
