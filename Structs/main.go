package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func (p person) printDetails() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func main() {

	kishan := person{
		firstName: "Kishan",
		lastName:  "Simha",
		contactInfo: contactInfo{
			email:   "kishan@kish.com",
			zipCode: 300,
		},
	}

	//kishanP := &kishan
	kishan.printDetails()
	kishan.updateName("Simhadriyaq")
	kishan.printDetails()

}
