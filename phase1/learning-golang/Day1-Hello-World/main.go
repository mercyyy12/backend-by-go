//Hello World and Variables

package main

import "fmt"

func main() {
	// Classic Hello World
	fmt.Println("Hello, World!")

	// Variable declarations with explicit type
	var name string
	name = "john"
	fmt.Println("Name:", name)

	// Short declaration with type inference
	name1 := "cena"
	fmt.Println("Name1:", name1)

	// Explicit type with initialization
	var num int = 20
	fmt.Println("Number:", num)

	// Implicit type with initialization
	var version = 13
	fmt.Println("Version:", version)

	// Short declaration for integers and floats
	num1 := 10
	num2 := 10.99
	fmt.Println("Num1:", num1)
	fmt.Println("Num2:", num2)

	// Boolean variable
	isTrue := true
	fmt.Println("IsTrue:", isTrue)

	// Constants
	const pi = 3.14
	fmt.Println("Pi:", pi)

	// Multiple variables declaration
	var a, b, c int = 1, 2, 3
	fmt.Println("Multiple vars:", a, b, c)

	// Multiple short declarations
	x, y, z := 4, 5, 6
	fmt.Println("Multiple short vars:", x, y, z)
}
