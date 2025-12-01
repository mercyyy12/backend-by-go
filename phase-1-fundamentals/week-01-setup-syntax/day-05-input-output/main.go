package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Basic printing
	name := "Mercyyy"
	fmt.Print(name)
	fmt.Println("\nMy name is ", name)
	fmt.Printf("My name is %v\n", name)

	// Scan input
	var a string
	var b int
	fmt.Println("Enter your name and age:")
	fmt.Scanln(&a, &b)
	fmt.Printf("Your name is %v and your age is %v\n", a, b)

	// Bufio Reader
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a sentence: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	text = strings.TrimSpace(text)
	fmt.Println("Reader read:", text)

	/*// Bufio Writer
	writer := bufio.NewWriter(os.Stdout)

	_, err := writer.WriteString("Hello i am Mercyyy!\n")
	if err != nil {
		fmt.Println("Error writing output:", err)
		return
	}

	// Always flush the buffer and handle error
	if err := writer.Flush(); err != nil {
		fmt.Println("Error flushing buffer:", err)
		return
	}

	// Scanner
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Type something (Ctrl + D to stop):")

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Scanner read:", line)
	}

	// After loop check for errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	} */

}
