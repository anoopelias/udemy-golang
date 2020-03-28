package main

import "fmt"

func main() {
	var colors map[string]string

	colors = make(map[string]string)

	colors["red"] = "#ff0000"
	colors["white"] = "#ffffff"

	printMap(colors)

}

func printMap(cs map[string]string) {

	for cs, h := range cs {
		fmt.Println("Hexcode for", cs, "is", h)
	}
}
