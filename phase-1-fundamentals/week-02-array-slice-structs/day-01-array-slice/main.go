package main

import "fmt"

func main() {
	// arrays (fixed size)
	var array [3]int // Array declaration
	array[0] = 7     //
	array[1] = 8
	fmt.Println(array, len(array)) // [7 8 0] 3, last element is 0 by default

	a := [7]int{0, 1, 2, 3, 4, 5, 6} // another way of array declaration
	fmt.Println(a, len(a))
	fmt.Printf("\n")

	// slice (dynamic size)
	slc := []int{0, 1, 2, 3, 4, 5, 6} // slice declaration
	fmt.Println("Length of slice is", len(slc), ", capacity of slice is", cap(slc))

	// nil and empty slice
	s := make([]int, 0)
	// var s []int // nil slice

	fmt.Println("Length of slice is", len(s), ", capacity of slice is", cap(s), s, s == nil)

	// slice using make()
	slice := make([]int, 2, 5) // slice declaration
	fmt.Printf("The length of slice is: %v\nThe capacity of the slice is: %v\nThe slice contains: %v\n\n", len(slice), cap(slice), slice)
	slice[0] = 100
	slice[1] = 200
	//slice[2] = 300 //error length = 2
	fmt.Println(slice) // [100 200]

	//append in slice
	slice = append(slice, 300)
	fmt.Printf("The length of slice is: %v\nThe capacity of the slice is: %v\nThe slice contains: %v\n\n", len(slice), cap(slice), slice)

	slice = append(slice, 400, 500, 600) //force growth, cap doubles
	fmt.Printf("The length of slice is: %v\nThe capacity of the slice is: %v\nThe slice contains: %v\n\n", len(slice), cap(slice), slice)

	slice2 := append(slice, 700)
	fmt.Printf("The length of slice is: %v\nThe capacity of the slice is: %v\nThe slice contains: %v\n\n", len(slice2), cap(slice2), slice2)

	// sub-slices
	slice = slice[2:]
	fmt.Println("sub slice is: ", slice)
	fmt.Println(len(slice), cap(slice))

	// memory sharing in slice
	destination := slice // destination points to the content of slice
	destination[0] = 7
	fmt.Printf("The value of slice is: %v\n\n", slice) //[7 400 500 600]

	// deep copy (making slices independent)
	dest := make([]int, len(slice), cap(slice))
	copy(dest, slice)
	dest = append(dest, 7, 8, 9, 1, 2, 3, 4, 5, 6)
	fmt.Println(slice, len(dest), cap(dest))

	// range loop in slice
	for i, v := range slice {
		fmt.Println(i, v)
	}
}
