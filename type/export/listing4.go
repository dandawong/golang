// Sample program to show how unexported fields from an exported
// struct type can't be accessed directly.
package main

import (
	"fmt"
	"github.com/dandawong/golang/type/export/entities"
)

func main() {
	// Create a value of type Admin from the entities package.
	a := entities.Admin{
		Rights: 10,
	}

	// Set the exported fields from the unexported inner type
	a.Name = "Bill"
	a.Email = "bill@email.com"

	fmt.Printf("User: %v\n", a)
}
