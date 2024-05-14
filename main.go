// main.go
package main

import (
	"fmt"
)

func main() {

	fmt.Println("Hello World")
	fmt.Println("-------------------------")

	// //! cannot use 89 (untyped int constant) as string value in variable declaration
	// // var name string = 89

	// var name string = "Arvind"
	// fmt.Println("Name:", name)

	// var version = "1.0"
	// fmt.Println("Version: " + version)

	// var money int = 67000
	// //! invalid operation: "Money:" + money (mismatched types untyped string and int)
	// // fmt.Println("Money:" + money)

	// fmt.Println("Money:", money)

	// var dimension float64 = 87.2
	// fmt.Println("dimension:", dimension)

	// var decided bool = false
	// fmt.Println("decided:", decided)

	// var person = "Arvind Chauhan"
	// fmt.Println("person:", person)

	// const pi = 3.14
	// fmt.Println("pi:", pi)

	// // pi = 3.15

	person := "Arvind Chauhan"
	fmt.Println("person:", person)

	// In Go, a variable, function, or type is considered public (or exported) if its name starts with an uppercase letter. Public entities can be accessed from outside the package where they are defined.

	Public := "data is important"
	private := "data is private"

	fmt.Println("Public:", Public)
	fmt.Println("private:", private)

}
