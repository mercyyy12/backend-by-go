/*
Challenge description:
Variables & types (string, int, bool)
User input with bufio.Reader
String to int conversion
Boolean logic (yes/no input)
Constants and iota
Formatted output with fmt.Printf
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Print("Enter your name: ")
	name := inputOutput()

	fmt.Print("Enter your age: ")
	s_age := inputOutput()
	//convrting string to int with error handling
	age, err := strconv.Atoi(s_age)
	if err != nil {
		fmt.Println("error", err)
		return
	}

	var activeStatus bool
	fmt.Print("Are you active? (yes or no): ")
	active := inputOutput()
	active = strings.ToLower(active)
	if active == "yes" {
		activeStatus = true
	} else {
		activeStatus = false
	}

	fmt.Printf("Your name is %v, age is %v and your active status is %v\n", name, age, activeStatus)

	const pi = 3.14
	const author = "Mercyyy"

	//iota example
	const (
		Beginner = iota
		Intermediate
		Advanced
	)
	fmt.Println("Pi:", pi)
	fmt.Println("Author:", author)
	fmt.Println("Levels:", Beginner, Intermediate, Advanced)

}

// get input from user using bufio reader
func inputOutput() string {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error", err)
		return "error"
	}
	text = strings.TrimSpace(text)
	return text
}
