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

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	jim.updateName()
	jim.print()
}

func (pp *person) updateName() {
	(*pp).firstName = "Jimmy"
}

func (pp person) print() {
	fmt.Printf("%+v", pp)
}
