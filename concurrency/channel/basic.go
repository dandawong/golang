package main

import (
	"fmt"
)

func main() {
	// Unbuffered channel of integers.
	unbuffered := make(chan int)

	// Buffered channel of strings.
	buffered := make(chan string, 10)

	// Sending values into a channel
	buffered <- "Gopher"

	// Receiving values from a channel
	value := <-buffered
}
