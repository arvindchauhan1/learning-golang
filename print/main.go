package main

import "fmt"

func main() {
	age := 21
	name := "Arvind Chauhan"
	points := 7.5123

	fmt.Println("Name:", name, "age:", age, "points:", points)
	fmt.Println("Hello World!")

	fmt.Printf("age is %d\n", age)
	// fmt.Printf("points are %f\n", points)
	fmt.Printf("points are %.3f\n", points)

	fmt.Printf("Type of name is %T\n", name)
	fmt.Printf("Type of name is %T\n", age)
	fmt.Printf("Type of name is %T\n", points)

}
