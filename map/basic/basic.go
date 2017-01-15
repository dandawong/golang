// https://play.golang.org/p/DaB_6Sa9Pt
package main

import (
	"fmt"
)

func main() {
	// Declaring a map using make
	map1 := make(map[string]int)
	map2 := map[string]string{"Red": "#da1337", "Orange": "#e95a22"}
	fmt.Println(map1)
	fmt.Println(map2)

	// Declaring an empty map using a map literal
	// panic
	// map3 := map[[]string]int{}
	// fmt.Println(map3)

	// Declaring a map that stores slices of strings
	map4 := map[int][]string{}
	fmt.Println(map4)

	// Assigning values to a map
	map5 := map[string]string{}
	map5["Red"] = "#da1337"
	fmt.Println(map5)

	// Runtime error assigned to a nil map
	var map6 map[string]string
	// panic
	// map6["Red"] = "#da1337"
	fmt.Println(map6)

	// Retrieving a value from a map and testing existence.
	map7 := map[string]string{"Red": "#da1337", "Orange": "#e95a22"}
	value, exists := map7["Red"]
	if exists {
		fmt.Println(value)
	}

	// Retrieving a value from a map testing the value for existence
	map8 := map[string]string{"Red": "#da1337", "Orange": "#e95a22"}
	value = map8["Blue"]
	// Only work if "" is not valid value of the map
	if value != "" {
		fmt.Println(value)
	}

	// Iterating over a map using for range
	map9 := map[string]string{
		"AliceBlue":   "#f0f8ff",
		"Coral":       "#ff7F50",
		"DarkGray":    "#a9a9a9",
		"ForestGreen": "#228b22",
	}
	for key, value := range map9 {
		fmt.Printf("Key: %s Value: %s\n", key, value)
	}

	// Removing an item from a map
	map10 := map[string]string{
		"AliceBlue":   "#f0f8ff",
		"Coral":       "#ff7F50",
		"DarkGray":    "#a9a9a9",
		"ForestGreen": "#228b22",
	}
	delete(map10, "Coral")
	for key, value := range map10 {
		fmt.Printf("Key: %s Value: %s\n", key, value)
	}

	// Passing maps between functions
	map11 := map[string]string{
		"AliceBlue":   "#f0f8ff",
		"Coral":       "#ff7F50",
		"DarkGray":    "#a9a9a9",
		"ForestGreen": "#228b22",
	}
	removeColor(map11, "Coral")
	for key, value := range map11 {
		fmt.Printf("Key: %s Value: %s\n", key, value)
	}
}

// removeColor removes keys from the specified map.
func removeColor(colors map[string]string, key string) {
	delete(colors, key)
}
