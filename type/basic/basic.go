// https://play.golang.org/p/LWHootXMPB
package main

import (
	"fmt"
)

func main() {
	// Declaration of a struct type
	type user struct {
		name       string
		email      string
		ext        int
		privileged bool
	}

	// Declaration of a variable of the struct type set to its zero value
	var bill user
	fmt.Println(bill)

	// Declaration of a variable of the struct type using a struct literal
	liza := user{
		name:       "Liza",
		email:      "lisa@email.com",
		ext:        123,
		privileged: true,
	}
	fmt.Println(liza)

	// Creating a struct type value without declaring the field names
	liza = user{"Lisa", "lisa@email.com", 456, true}
	fmt.Println(liza)

	// Declaring fields based on other struct types
	type admin struct {
		person user
		level  string
	}

	// Using struct literals to create values for fields
	fred := admin{
		person: user{
			name:       "Fred",
			email:      "fred@email.com",
			ext:        111,
			privileged: true,
		},
		level: "super",
	}
	fmt.Println(fred)

	// Declaration of a new type based on an int64
	type Duration int64

	// Compiler error assigning value of different types
	var dur Duration
	// panic
	// dur = int64(1000)
	fmt.Println(dur)
}
