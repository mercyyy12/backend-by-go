package main

import (
	"fmt"
	"strings"
)

func main() {

	// fixed-length array declaration
	arr := [3]int{10, 20, 30}
	fmt.Println("Array:", arr)

	// slice with make (length 2, capacity 5)
	sl := make([]int, 2, 5)
	sl[0] = 1
	sl[1] = 2
	fmt.Println("Initial slice:", sl)

	// Copy a slice
	original := []int{1, 2, 3, 4, 5}
	copied := make([]int, len(original))
	copy(copied, original)
	copied[0] = 99
	fmt.Println("Original slice:", original)
	fmt.Println("Copied slice:", copied)

	// Slicing a slice
	sub := original[1:4] // elements at index 1 to 3
	fmt.Printf("len: %d cap: %d\n", len(sub), cap(sub))

	// Strings are immutable
	str := "hello"
	fmt.Println("Original string:", str)
	// str[0] = 'H' // not allowed

	// Common string operations
	fmt.Println("ToUpper:", strings.ToUpper("go"))            // "GO"
	fmt.Println("Contains:", strings.Contains("hello", "he")) // true
	fmt.Println("Split:", strings.Split("a,b,c", ","))        // [a b c]
	fmt.Println("TrimSpace:", strings.TrimSpace("  hello  ")) // "hello"

}
