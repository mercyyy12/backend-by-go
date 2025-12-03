package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// initialized otp
	const otp = 123456
	chance := 4

outer:
	for {
		// input from user
		text := io()
		if text == -1 {
			fmt.Println("Enter valid number")
			text = io()
		}

		// if otp is correct
		if text == otp {
			fmt.Println("successfully entered!")
			break outer
		}
		chance--
		var j int
		switch chance {
		case 3:
			j = 1
		case 2:
			j = 3
		case 1:
			j = 6
		}
		fmt.Printf("Wrong otp. Please wait\nYou have %v attempts remaining\n", chance)

		// progressive lockout
		for i := 5 * j; i > 0; i-- {
			fmt.Printf("\r%d... ", i)
			time.Sleep(1 * time.Second)
		}
		fmt.Printf("\n")
		if chance <= 0 {
			println("Permanently locked")
			break outer
		}
	}
}

func io() int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the OTP")
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading the number", err)
		return -1
	}

	// converting string to int
	text = strings.TrimSpace(text)
	ftext, err := strconv.Atoi(text)
	if err != nil {
		fmt.Println("Error reading the number", err)
		return -1
	}

	// fmt.Printf("the value is %v and type is %T", ftext, ftext)
	return ftext
}
