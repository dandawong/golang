// https://play.golang.org/p/kLt5O-yR3H
// Sample program to show how you can't always get the
// address of a value
package main

import "fmt"

// duration is a type with a base type of int.
type duration int

// format pretty-prints the duration value.
func (d *duration) pretty() string {
	return fmt.Sprintf("Duration: %d", *d)
}

func main() {
	duration(42).pretty()
	// cannot call pointer method on duration(42)
	// cannot take the address of duration(42)

	d := duration(42)
	fmt.Println(d)
	fmt.Printf("%p\n", &d)
	fmt.Println(d.pretty())

	// panic
	duration(33).pretty()
}
