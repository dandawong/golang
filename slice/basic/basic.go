// https://play.golang.org/p/LM3uXMy729
package main

import (
	"fmt"
)

func main() {
	// Declaring a slice of strings by length
	slice1 := make([]string, 5)
	fmt.Println(slice1)

	// Declaring a slice of integers by length and capacity
	slice2 := make([]int, 3, 5)
	fmt.Println(slice2)
	// panic
	// fmt.Println(slice2[3])

	// Compiler error setting capacity less than length
	// panic
	// slice3 := make([]int, 5, 3)

	// Declaring a slice with a slice literal, no ... inside []
	slice4 := []string{"a", "b", "c", "d", "e"}
	fmt.Println(slice4)
	slice5 := []int{1, 2, 3}
	fmt.Println(slice5)

	// Declaring a slice with index positions
	slice6 := []string{99: "abc"}
	fmt.Println(slice6)

	// Declaration differences between arrays and slices
	array1 := [3]int{4, 5, 6}
	slice7 := []int{7, 8, 9}
	fmt.Println(array1)
	fmt.Println(slice7)

	// Declaring a nil slice
	var slice8 []int
	fmt.Printf("%p\n", &slice8)
	slice8 = append(slice8, 1)
	fmt.Println(slice8)
	fmt.Printf("%p\n", &slice8)

	// Declaring an empty slice
	slice9 := make([]int, 0)
	slice10 := []int{}
	fmt.Println(slice9)
	fmt.Println(slice10)

	// Declaring an array using an array literal
	slice11 := []int{10, 20, 30, 40, 50}
	slice11[1] = 25
	fmt.Println(slice11)

	// Taking the slice of a slice
	slice12 := []int{10, 20, 30, 40, 50}
	slice13 := slice12[1:3]
	fmt.Println(slice13)

	// Potential consequence of making changes to a slice
	slice14 := []int{10, 20, 30, 40, 50}
	slice15 := slice14[1:3]
	slice15[1] = 35
	fmt.Println(slice14)

	// Runtime error showing index out of range
	slice16 := []int{10, 20, 30, 40, 50}
	slice17 := slice16[1:3]
	fmt.Println(slice17)
	// panic
	// slice17[2] = 45

	// Using append to add an element to a slice
	slice18 := []int{10, 20, 30, 40, 50}
	slice19 := slice18[1:3]
	slice19 = append(slice19, 60)
	fmt.Println(slice18)
	fmt.Println(slice19)
	for i := 0; i < 3; i++ {
		slice19 = append(slice19, (i+1)*10)
	}
	// two slices not point to same array
	fmt.Println(slice18)
	fmt.Println(slice19)
	slice19[2] = 55
	fmt.Println(slice18)
	fmt.Println(slice19)

	// Using append to increase the length and capacity of a slice
	slice20 := []int{10, 20, 30, 40}
	slice21 := append(slice20, 50)
	slice21[0] = 1
	fmt.Println(slice20)
	fmt.Println(slice21)

	// Declaring a slice of string using a slice literal
	slice22 := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
	slice23 := slice22[2:3:4]
	// panic
	// slice24 := slice22[2:3:6]
	fmt.Println(slice23)

	// Benefits of setting length and capacity to be the same
	slice25 := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
	slice26 := slice25[2:3:3]
	slice26 = append(slice26, "Kiwi")
	fmt.Println(slice25)
	fmt.Println(slice26)

	// Appending to a slice from another slice
	slice27 := []int{1, 2}
	slice28 := []int{3, 4}
	slice29 := append(slice27, slice28...)
	fmt.Println(slice27)
	fmt.Println(slice29)

	// Iterating over a slice using for range
	slice30 := []int{10, 20, 30, 40}
	for index, value := range slice30 {
		fmt.Printf("Index: %d Value: %d\n", index, value)
	}

	// range provides a copy of each element
	slice31 := []int{10, 20, 30, 40}
	for index, value := range slice31 {
		fmt.Printf("Value: %d Value-addr: %X ElemAddr: %X\n", value, &value, &slice31[index])
	}

	// Iterating over a slice using a traditional for loop
	slice32 := []int{10, 20, 30, 40}
	for index := 2; index < len(slice32); index++ {
		fmt.Printf("Index: %d Value: %d\n", index, slice32[index])
	}

	// Declaring a multidimensional slice
	slice33 := [][]int{{10}, {100, 200}}
	fmt.Println(slice33)

	// Composing slices of slices
	slice34 := [][]int{{10}, {100, 200}}
	slice34[0] = append(slice34[0], 20)
	fmt.Println(slice34)
	slice34[1] = append(slice34[0], 30)
	fmt.Println(slice34)
	slice34[1][0] = 9
	fmt.Println(slice34)

	// Passing slices between functions
	slice35 := make([]int, 1e6)
	slice35 = foo(slice35)
}

// Function foo accepts a slice of integers and returns the slice back.
func foo(slice []int) []int {
	return slice
}
