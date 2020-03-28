package main

import "fmt"

type bot interface {
	getGreeting() string
}

type engishBot struct{}
type spanishBot struct{}

func main() {
	e := engishBot{}
	s := spanishBot{}

	printGreeting(e)
	printGreeting(s)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (engishBot) getGreeting() string {
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
