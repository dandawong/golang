// Sample program to show how unexported fields from an exported
// struct type can't be accessed directly.
package main

import (
	"fmt"
	"github.com/dandawong/golang/type/export/entities"
)

func main() {
	u := entities.User{
		Name:  "Bill",
		email: "bill@email.com",
	}
	// unknown entities.User field 'email' in struct literal
	fmt.Printf("User: %v\n", u)
}
