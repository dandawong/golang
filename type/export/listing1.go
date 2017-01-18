// Sample program to show how the program can't access an
// unexported identifier from another package.
package main

import (
	"fmt"
	"github.com/dandawong/golang/type/export/counters"
)

func main() {
	// Create a variable of the unexported type and initialize
	// the value to 10.
	counter := counters.alertCounter(10)

	// cannot refer to unexported name counters.alertCounter
	// undefined: counters.alertCounter
	fmt.Printf("Counter: %d\n", counter)
}
