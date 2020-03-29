package main

import (
	"fmt"
	"net/http"
)

func main() {

	links := []string{
		"http://www.google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://amazon.com",
		"http://stackoverflow.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)

	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "Might be down")
		c <- "Must be down"
		return
	}

	fmt.Println(link, "is up!")
	c <- "It is up"

}
