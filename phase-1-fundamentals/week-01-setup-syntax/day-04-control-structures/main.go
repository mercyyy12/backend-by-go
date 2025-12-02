package main

import "fmt"

func main() {
	// if structure
	// 1. with condition only
	x := 10
	y := 21
	if x > 5 {
		fmt.Println("x is greater than 5")
	}

	// 2. with init, condition and else
	if y := 0; y >= 0 {
		fmt.Println("Number is positive")
	} else {
		fmt.Println("Number is negative")
	}

	// 3. with init, condition and else
	if x%2 == 0 && y%2 == 0 {
		fmt.Println("x and y both are even numbers")
	} else if x < 0 && y < 0 {
		fmt.Println("x and y are negative numbers")
	} else {
		fmt.Printf("The value of x is %v and y is %v\n", x, y)
	}

	// switch
	// 1. with expression
	day := 5

	switch day {
	case 1:
		fmt.Println("its sunday")
	case 2:
		fmt.Println("monday")
	case 3:
		fmt.Println("tuesday")
	default:
		fmt.Println("invalid")
	}

	//2. Multiple cases in one line
	var ext string = "jpg"
	switch ext {
	case "jpg", "jpeg", "png", "gif":
		fmt.Println("Image file")
	}

	//3. Fallthrough
	switch n := 2; n {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
		fallthrough
	case 3:
		fmt.Println("Three")
	}

	// 4. switch without expression
	score := 82
	switch {
	case score >= 90:
		fmt.Println("A")
	case score >= 80:
		fmt.Println("B")
	default:
		fmt.Println("F")
	}

}
