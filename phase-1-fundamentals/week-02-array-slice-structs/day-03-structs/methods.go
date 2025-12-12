package main

import "fmt"

type Person struct {
	Name  string
	phone int
	isBirthday
}

type isBirthday struct {
	Age int
}

// Value receiver method
func (p Person) Greet() {
	fmt.Println("hello", p.Name, p.phone)
}

// Pointer receiver method
func (st *Person) birthday() {
	st.Age++
}
