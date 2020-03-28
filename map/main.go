package main

import "fmt"

func main() {
	var colors map[string]string

	colors = make(map[string]string)

	colors["red"] = "#ff0000"
	colors["white"] = "#ffffff"

	fmt.Println(colors)

}
