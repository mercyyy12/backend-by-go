package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("enter your name: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("error", err)
		return
	}
	name := strings.TrimSpace(input)
	// name, err := strconv.Atoi(input)
	// if err != nil {
	// 	fmt.Println("error", err)
	// 	return
	// }
	fmt.Printf("your name from reader is %s\n", name)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your name: ")
	if !scanner.Scan() {
		fmt.Println("Failed to read input.")
		return
	}
	name1 := strings.TrimSpace(scanner.Text())
	// age, err := strconv.Atoi(input1)
	// if err != nil {
	// 	fmt.Println("Invalid age")
	// 	return
	// }
	fmt.Println("You are", name1, "years old.")
}
