package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Arg missing : file name")
		os.Exit(1)
	}

	filename := os.Args[1]
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("Unable to read file", err)
		os.Exit(2)
	}

	io.Copy(os.Stdout, file)
}
