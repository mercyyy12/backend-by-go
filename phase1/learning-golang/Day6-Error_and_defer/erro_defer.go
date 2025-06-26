package main

import (
	"errors"
	"fmt"
)

// posneg returns an error if number is negative
func posneg(x int) error {
	if x < 0 {
		return errors.New("is negative")
	}
	return nil
}

// div returns a / b and an error if b is zero
func div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("divisor cannot be 0")
	}
	return a / b, nil
}

func main() {

	fmt.Print("Enter a number to check positive/negative: ")
	var num int
	if _, err := fmt.Scanf("%d\n", &num); err != nil {
		fmt.Println("input error:", err)
		return
	}

	if err := posneg(num); err != nil { // err exists only in this if-block
		fmt.Printf("error: %s\n", err)
		return
	}
	fmt.Println("positive number")

	//err reused after the if -------------
	fmt.Println("Division")
	var a, b int
	fmt.Print("Enter dividend and divisor: ")
	if _, err := fmt.Scanf("%d %d\n", &a, &b); err != nil {
		fmt.Println("input error:", err)
		return
	}

	result, err := div(a, b) // err stays in scope after the if
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("result:", result)
}
