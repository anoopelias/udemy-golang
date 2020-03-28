package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)

	var bruce person
	bruce.firstName = "Bruce"
	bruce.lastName = "Banner"

	fmt.Println(bruce)
	fmt.Printf("%+v", bruce)

}
