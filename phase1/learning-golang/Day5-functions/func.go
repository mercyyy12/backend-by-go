package main

import "fmt"

func main() {
	// Call greet function (no parameters, no return)
	greet()

	// Call greet1 with a string parameter
	greet1("santosssh")

	// Call add function with two integers
	fmt.Println("add:", add(5, 10))

	// Call squareAndCube which returns two values square and cube of x
	sq, cb := squareAndCube(2)
	fmt.Printf("square = %d, cube = %d\n", sq, cb)
}

// greet prints a simple hello message
func greet() {
	fmt.Println("Hello world!")
}

// greet1 prints hello message with the provided name
func greet1(name string) {
	fmt.Println("Hello ", name)
}

// add takes two integers and returns their sum
func add(a, b int) int {
	return a + b
}

// squareAndCube calculates and returns the square and cube of the input x. uses named return values: square and cube
func squareAndCube(x int) (square, cube int) {
	square = x * x
	cube = x * x * x
	return // returns square and cube implicitly
}
