package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	Embeded_structure()

	// creating struct instance
	p1 := person{"Mercyyy", 25}
	// or p1 := person{name: "Mercyyy", age: 25}
	fmt.Println(p1)
	/*
		// zero value initialization
		var p person
		p.name = "NoMercyyy" // accessing or updating field
		p.age = 20
		fmt.Println(p)
	*/

	/*
		// Using new
		p := new(person)
		p.name = "Mercyyy"
		p.age = 20
		fmt.Println(p)
	*/
	// using & literal
	p := &person{"Legion", 2024}
	fmt.Println(p)

	// copying struct (except for pointers) = only values gets copied
	p2 := p1
	p2.name = "Mercyyy2"
	println("the name of person 1 is: ", p1.name)
	println("the name of person 2 is: ", p2.name)

}
