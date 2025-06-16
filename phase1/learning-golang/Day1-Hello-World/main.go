package main

import (
	"fmt"
)

func main() {
	// Classic Hello World
	fmt.Println("Hello, World!") //prints the text to the console with a newline at the end

	// Variables
	// Explicit type declaration
	var name string
	name = "John"
	fmt.Println("Name:", name)

	// Declaring and initializing with explicit type
	var age int = 20
	fmt.Println("Age:", age)

	// Short declaration with type inference
	username := "Cena"
	fmt.Println("Username:", username)

	// Implicit type
	var version = 1.21
	fmt.Println("Version:", version)

	// Short declaration for int and float
	num1 := 42
	num2 := 3.14
	fmt.Println("Num1:", num1, "Num2:", num2)

	// Boolean variable
	isActive := true
	fmt.Println("IsActive:", isActive)

	//Constants
	const pi = 3.14159
	const author = "Santosh"
	fmt.Println("Pi:", pi)
	fmt.Println("Author:", author)

	//Multiple Declarations explicit and implicit
	var a, b, c int = 1, 2, 3
	fmt.Println("Multiple vars:", a, b, c)

	x, y, z := "backend", "with", "go"
	fmt.Println("Multiple vars:", x, y, z)

	// Zero values
	var defaultInt int
	var defaultString string
	var defaultBool bool

	fmt.Println("Default int:", defaultInt)       // 0
	fmt.Println("Default string:", defaultString) // ""
	fmt.Println("Default bool:", defaultBool)     // false

}
