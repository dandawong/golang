// https://play.golang.org/p/rRDXb3tGRK
package main

import (
	"fmt"
)

func main() {
	// Declare an integer array of five elements.
	var array1 [5]int
	fmt.Println(array1)

	// Declaring an array using an array literal
	array2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array2)

	// Declaring an array with Go calculating size
	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	// Declaring an array initializing specific elements
	array4 := [5]int{1: 1, 3: 4}
	fmt.Println(array4)

	// Accessing array elements
	array5 := [5]int{10, 20, 30, 40, 50}
	array5[3] = 110
	fmt.Println(array5)

	// Accessing array pointer elements
	var array6 = [5]*int{1: new(int), 3: new(int)}
	fmt.Println(array6)
	*array6[1] = 34
	*array6[3] = 128
	fmt.Println(*array6[1])
	// panic
	// fmt.Println(*array6[2])

	// Assigning one array to another of the same type
	// panic
	// var array7 [6]string
	var array7 [5]string
	array8 := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}
	array7 = array8
	fmt.Println(array7)

	// Assigning one array of pointers to another
	var array9 [3]*string
	array10 := [3]*string{new(string), new(string), new(string)}
	*array10[0] = "Red"
	*array10[1] = "Blue"
	*array10[2] = "Green"
	array9 = array10
	for _, p := range array9 {
		fmt.Println(*p)
	}

	// Declaring two-dimensional arrays
	var array11 [4][2]int
	array12 := [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	array13 := [4][2]int{1: {20, 21}, 3: {40, 41}}
	array14 := [4][2]int{1: {0: 20}, 3: {1: 41}}
	fmt.Println(array11)
	fmt.Println(array12)
	fmt.Println(array13)
	fmt.Println(array14)

	// Accessing elements of a multi-dimensional array
	var array15 [2][2][2]int
	array15[1][0][0] = 10
	array15[1][0][1] = 20
	array15[1][1][0] = 30
	array15[1][1][1] = 40
	fmt.Println(array15)

	// Assigning multidimensional arrays by index
	array16 := [2][2]int{{1, 2}, {3, 4}}
	var array17 [2]int = array16[1]
	var value1 int = array16[0][1]
	fmt.Println(array17)
	fmt.Println(value1)

	// Passing array between function
	var array18 [1e6]int
	foo(array18)
	bar(&array18)
}

// Function foo accepts an array of one million integers.
func foo(array [1e6]int) {

}

// Function bar accepts a pointer to an array of one million integers.
func bar(array *[1e6]int) {

}
