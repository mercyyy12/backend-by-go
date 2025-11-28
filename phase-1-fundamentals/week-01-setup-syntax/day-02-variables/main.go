package main

import "fmt"

func main() {
	// Variables

	// Implicit type
	var version = 1.21
	fmt.Println("Version:", version)
	//Gives type and value of version as OP
	fmt.Printf("The value of version is %v, and it's type is %T\n", version, version)

	// Explicit type declaration
	var name string
	name = "John"
	fmt.Println("Name:", name)
	//	Gives type and value of name as OP
	fmt.Printf("The value of name is %v, and it's type is %T\n", name, name)

	// Declaring and initializing with explicit type
	var age int = 20
	fmt.Println("Age:", age)

	// Short declaration or type inference
	username := "Cena"
	fmt.Println("Username:", username)

	// Short declaration or type inference for int and float
	num1 := 42
	num2 := 3.14
	fmt.Println("Num1:", num1, "Num2:", num2)

	// Boolean variable
	isActive := true
	fmt.Println("IsActive:", isActive)
	//	Gives type and value of boolean value as OP
	fmt.Printf("The value of isActive is %v, and it's type is %T\n", isActive, isActive)

	//Constants
	const pi = 3.14159
	const author = "Mercyyy"
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

	//iota
	const (
		f = iota //0
		g        //1
		h        //2
	)
	fmt.Println(f, g, h)
}
