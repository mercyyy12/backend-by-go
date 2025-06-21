package main

import "fmt"

func main() {

	// // fixed-length array declaration
	// arr := [3]int{10, 20, 30}
	// fmt.Println("Array:", arr)

	// // slice with make (length 2, capacity 5)
	// sl := make([]int, 2, 5)
	// sl[0] = 1
	// sl[1] = 2
	// fmt.Println("Initial slice:", sl)

	// // Copy a slice
	// original := []int{1, 2, 3, 4, 5}
	// copied := make([]int, len(original))
	// copy(copied, original)
	// copied[0] = 99
	// fmt.Println("Original slice:", original)
	// fmt.Println("Copied slice:", copied)

	// // Slicing a slice
	// sub := original[1:4] // elements at index 1 to 3
	// fmt.Printf("len: %d cap: %d\n", len(sub), cap(sub))

	// // Strings are immutable
	// str := "hello"
	// fmt.Println("Original string:", str)
	// // str[0] = 'H' // not allowed

	// // Common string operations
	// fmt.Println("ToUpper:", strings.ToUpper("go"))            // "GO"
	// fmt.Println("Contains:", strings.Contains("hello", "he")) // true
	// fmt.Println("Split:", strings.Split("a,b,c", ","))        // [a b c]
	// fmt.Println("TrimSpace:", strings.TrimSpace("  hello  ")) // "hello"

	//maps
	id := map[string]string{
		"santosh": "7412",
		"ram":     "645",
		"pemsang": "54545",
	}
	id["shyam"] = "45"
	delete(id, "ram")
	for k, v := range id {
		fmt.Println(k, "is", v)
	}

	// Range with Slice
	numbers := []int{10, 20, 30, 40}
	fmt.Println("Slice:")
	for i, val := range numbers {
		fmt.Printf("Index %d: Value %d\n", i, val)
	}

	// Range with Array
	arr := [3]string{"go", "is", "cool"}
	fmt.Println("Array:")
	for i, word := range arr {
		fmt.Printf("Index %d: Value %s\n", i, word)
	}

	// Range with String
	text := "Go"
	fmt.Println("String:")
	for i, ch := range text {
		fmt.Printf("Byte Index %d: Rune %c\n", i, ch)
	}

	// Range with Map
	ages := map[string]int{
		"Alice": 25,
		"Bob":   30,
		"Eve":   22,
	}
	fmt.Println("Map:")
	for name, age := range ages {
		fmt.Printf("%s is %d years old\n", name, age)
	}
}
