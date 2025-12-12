package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	// creating struct instance
	p := person{"Mercyyy", 25}
	// or p1 := person{name: "Mercyyy", age: 25}
	fmt.Println("Creation of basic struct instance: ", p)
	fmt.Printf("\n")

	// zero value initialization
	var p1 person
	p1.name = "NoMercyyy" // accessing or updating field
	p1.age = 20
	fmt.Println("An example of Zero value initialization of struct: ", p1)
	fmt.Printf("\n")

	// Using new
	p2 := new(person)
	p2.name = "Mercyyy"
	p2.age = 20
	fmt.Println("Creation of struct using 'new' literal: ", p2)
	fmt.Printf("\n")

	// using & literal
	p3 := &person{"Legion", 2024}
	fmt.Println("Creation of struct using '&' literal: ", p3)
	fmt.Printf("\n")

	// copying struct (except for pointers) = only values gets copied
	fmt.Println("An example to showcase the struct copy")
	p4 := p1
	p4.name = "Mercyyy2"
	println("the name of person 1 is: ", p1.name)
	println("the name of person 2 is: ", p4.name)
	fmt.Printf("\n")

	// JSON tag call
	fmt.Println("Example of JSON tags: ")
	Json_tags()
	fmt.Printf("\n")

	// Methods call
	fmt.Println("Example of Methods: ")
	p5 := Person{"Mercyyy", 25, isBirthday{20}}
	p5.Greet()
	p5.birthday()
	fmt.Println(p5)
	fmt.Printf("\n")

	// Embeded-structure call
	fmt.Println("An example of Embeded-structure: ")
	Embeded_structure()
	fmt.Printf("\n")

}
