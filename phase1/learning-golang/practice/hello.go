package main

import (
	"fmt"
)

func main() {
	// Range over Slice
	fmt.Println(" Slice:")
	numbers := []int{10, 20, 30, 40}
	for index, value := range numbers {
		fmt.Printf("Index %d: Value %d\n", index, value)
	}

	//  Range over String
	fmt.Println("\n String (Unicode-safe):")
	text := "Go💙Lang"
	for i, ch := range text {
		fmt.Printf("Index %d: Rune %q\n", i, ch)
	}

	//Range over Map
	fmt.Println(" Map:")
	person := map[string]string{
		"name": "Santosh",
		"city": "Kathmandu",
	}
	for key, value := range person {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}
}
