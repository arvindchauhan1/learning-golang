// since packages are tightly coupled with folders

// In Go, the convention is to have a one-to-one relationship between folders and packages. This means that each folder typically corresponds to a single package.

package myutil

import (
	"fmt"
)

func PrintMessage(message string) {
	fmt.Println(message)
	
}
