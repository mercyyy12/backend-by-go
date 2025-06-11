package main

import (
	"fmt"
	"strings"
)

func main() {
	//for loop w/o clauses is used for creating infinite loop
	for {
		var choice int
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Check Even or Odd")
		fmt.Println("2. Print Multiplication Table")
		fmt.Println("3. Check Leap Year")
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			EvenOdd()
		case 2:
			multiplication()
		case 3:
			checkLeapYear()
		default:
			fmt.Println("Enter a valid option (1, 2, or 3).")
		}

		var decision string
		fmt.Print("\nDo you want to try again? (y/n): ")
		fmt.Scan(&decision)
		decision = strings.ToLower(strings.TrimSpace(decision))

		//decides whether to continue loop or break it
		if decision != "y" {
			break
		}
	}
}

// Check if number is even or odd using if else
func EvenOdd() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	if num%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}
}

// Print multiplication table using for
func multiplication() {
	var num int
	fmt.Print("Enter a number for multiplication: ")
	fmt.Scan(&num)

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", num, i, num*i)
	}
}

// Check if a year is a leap year using if else if
func checkLeapYear() {
	var year int
	fmt.Print("Enter a year: ")
	fmt.Scan(&year)

	if year%400 == 0 {
		fmt.Println("Leap year")
	} else if year%100 == 0 {
		fmt.Println("Not a leap year")
	} else if year%4 == 0 {
		fmt.Println("Leap year")
	} else {
		fmt.Println("Not a leap year")
	}

}
