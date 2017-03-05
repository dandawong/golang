package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go hello(ch)
	for {
		v, ok := <-ch
		if ok {
			fmt.Println(v)
		} else {
			break
		}
	}
}

func hello(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}
