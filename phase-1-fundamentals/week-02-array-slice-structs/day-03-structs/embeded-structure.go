package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type student struct {
	name string
	age  int

	// Embeded
	Address
}

type Address struct {
	temp_address string
	phone_number int
}

func Embeded_structure() {

	// store multiple records in structure

	var n int
	fmt.Println("Enter no. of students")
	fmt.Scanf("%d\n", &n)

	// init slice to store data of students
	p := make([]student, n)

	// initialized bufio for string input
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < n; i++ {
		fmt.Printf("Enter name of student %v: ", i+1)
		name, _ := reader.ReadString('\n')
		p[i].name = strings.TrimSpace(name)

		fmt.Printf("Enter age: ")
		fmt.Scanf("%d", &p[i].age)

		fmt.Printf("Enter address: ")
		address, _ := reader.ReadString('\n')
		p[i].temp_address = strings.TrimSpace(address)

		fmt.Printf("Enter phone number: ")
		fmt.Scanf("%d", &p[i].phone_number)
	}

	// to display the data using range
	for i, s := range p {
		fmt.Printf("The name of student %v = %s\nAge = %v\nTemporary address = %v\nPhone number = %v\n",
			i+1, s.name, s.age, s.temp_address, s.phone_number)
	}
}
