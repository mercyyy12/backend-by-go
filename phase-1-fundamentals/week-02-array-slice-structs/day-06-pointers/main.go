package main

import "fmt"

type student struct {
	name string
	age  int
}

func main() {
	x := 10
	fmt.Println("The value of x is:", x)

	// & pointer stores the memory address of a value
	fmt.Println("The memmory address of x is:", &x) // memory address

	y := &x // address of x is stored in y
	// * dereferences the address
	fmt.Println("The value of y is:", *y) //10\
	age := 20

	passValue(age)
	fmt.Println("Your age using pass by value is:", age)
	passPointer(&age)
	fmt.Println("Your age using pass by reference is:", age)

	var s student
	s.name = "Neymar"
	fmt.Println("Without pointer func, the name of student is:", s.name)
	changeName(&s)
	fmt.Println("After pointer func, the name of student is:", s.name)

	s.age = 20
	s.ageInc()
	fmt.Println("After pointer func, the age of student is:", s.age)
}

// Pass by value (no change in original data)
func passValue(age int) {
	age = age + 5
	fmt.Println("Your age using pass by value is:", age)
}

// Pass by pointer (change in original data)
func passPointer(age *int) {
	*age++
}

// Pointer in struct
func changeName(s *student) {
	s.name = "Mercyyy"
}

// Pointer receiver in methods
func (s *student) ageInc() {
	s.age++
}
