package main

import (
	"errors"
	"fmt"
)

// Initializer sentinel error
var ErrDivisionByZero = errors.New("division by zero")

func main() {

	// var err error = nil
	// if err != nil {
	// 	fmt.Println("There was an error:", err)
	// } else {
	// 	fmt.Println("No error!")
	// }

	a, b := 10, 0
	c, err := divide(a, b)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println(c)
}

// Returning error from function
func divide(a, b int) (int, error) {
	if b == 0 {

		//panic stops prog immediately
		// panic("Can't divide by zero!")

		// Creating simple error
		// return 0, errors.New("cannot divide by zero!")

		// Creating fromatted error
		//return 0, fmt.Errorf("Can't divide %d by %d", a, b)

		// Sentinel error
		return 0, ErrDivisionByZero
	}
	return a / b, nil
}
