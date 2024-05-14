package main

import "fmt"

func simpleFunction() {
	fmt.Println("This is a simple function")
}

// func add(a int, b int) int {
// 	return a + b
// }

// named return value
func add(a int, b int) (result int) {
	result = a + b
	return
}

func main() {
	fmt.Println("Learnig about function in GoLang")
	simpleFunction()

	ans := add(2, 3)
	fmt.Println("Addition of 2 and 3 is:", ans)
}
