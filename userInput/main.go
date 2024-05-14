package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hey,What's your name?")

	//! fmt.Scanln()
	// it'll read only one word
	var name string
	fmt.Scanln(&name)
	fmt.Println("Hello Mr.", name)

	//! bufio.NewReader()
	// it'll read the whole line

	fmt.Println("What's your fullname ?")

	// creates a new reader whose input is from the standard input
	// os.Stin read from os standard input
	reader := bufio.NewReader(os.Stdin)

	// read the whole line till the new line character ('\n')
	name, _ = reader.ReadString('\n')
	fmt.Println("Hello Mr.", name)
}
