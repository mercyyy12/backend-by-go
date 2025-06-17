package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Uncomment the section you want to test.

	// fmtBasicInput()
	// bufioReaderInput()
	// bufioFileReader()
	// bufioScannerInput()
	// bufioFileScanner()
	// bufioWriterExample()
}

// fmt package basic input examples
func fmtBasicInput() {
	var age int
	var health, name string

	fmt.Print("Enter your age: ") // prints text without newline
	fmt.Scan(&age)                // reads input until space or newline

	fmt.Println("How are you feeling today?") // prints text with newline
	fmt.Scanln(&health)                       // reads input until newline

	fmt.Print("What is your name? ")
	fmt.Scanln(&name)

	fmt.Printf("Hello %s, your age is %d and you're feeling %s today.\n", name, age, health)

	// Example of fmt.Scanf with multiple inputs
	var day int
	var month string
	fmt.Print("Enter your birthday day and month: ")
	fmt.Scanf("%d %s", &day, &month)
	fmt.Printf("Your birthday is on %d of %s.\n", day, month)
}

// bufio.Reader from stdin
func bufioReaderInput() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your age: ")
	input, err := reader.ReadString('\n') // reads until newline
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	input = strings.TrimSpace(input) // clean newline/spaces
	age, err := strconv.Atoi(input)  // convert string to int
	if err != nil {
		fmt.Println("Invalid input")
		return
	}
	fmt.Println("You are", age, "years old.")
}

// bufio.Reader reading from file
func bufioFileReader() {
	file, err := os.Open("data.txt") // open file data.txt file must exist in the same directory.
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	fmt.Println("Reading file line by line:")
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Print(line)
	}
}

// bufio.Scanner from stdin
func bufioScannerInput() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your age: ")
	if !scanner.Scan() {
		fmt.Println("Failed to read input.")
		return
	}
	input := strings.TrimSpace(scanner.Text())
	age, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid age")
		return
	}! Please enter a number.
	fmt.Println("You are", age, "years old.")
}

// bufio.Scanner reading from file
func bufioFileScanner() {
	file, err := os.Open("data.txt") //data.txt file must exist in the same directory.
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fmt.Println("File content line by line:")
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Scan error:", err)
	}
}

// bufio.Writer writing to file
func bufioWriterExample() {
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	lines := []string{
		"Line 1: Hello, World!\n",
		"Line 2: This is Go.\n",
		"Line 3: Writing to a file.\n",
	}

	for _, line := range lines {
		writer.WriteString(line)
	}
	writer.Flush()
	fmt.Println("Data written to output.txt successfully.")
}
