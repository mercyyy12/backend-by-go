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
	//converting string to int with error handling
	age, err := strconv.Atoi(s_age)
	if err != nil {
		fmt.Println("error", err)
		return
	}

	fmt.Print("Enter your salary: ")
	salaries := inputOutput()
	salary, err := strconv.ParseFloat(salaries, 64)
	if err != nil {
		fmt.Println("Invalid salary!")
		return
	}

	remoteW := false
	fmt.Print("Are you remote worker? (yes/no): ")
	remote := inputOutput()
	remote = strings.ToLower(remote)
	if remote == "yes" {
		remoteW = true
	} else {
		remoteW = false
	}
	fmt.Println("\nProfile card")
	fmt.Println("_______________________________")
	fmt.Printf("Name\t: %v\t\t%T\n", name, name)
	fmt.Printf("Age\t: %v\t\t%T\n", age, age)
	fmt.Printf("Salary\t: %v\t\t%T\n", salary, salary)
	fmt.Printf("Remote\t: %v\t\t%T\n", remoteW, remoteW)

	// iota example
	fmt.Println("\nUser Role")
	fmt.Println("_______________________________")
	const (
		Admin = iota
		Editor
		Viewer
	)
	fmt.Printf("%v= Admin\n%v= Editor\n%v= Viewer\n", Admin, Editor, Viewer)
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
