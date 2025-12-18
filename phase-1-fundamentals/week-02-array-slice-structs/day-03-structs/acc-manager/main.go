package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Map init
var m = make(map[any]*User)

// user structure
type User struct {
	Name    string
	Number  string
	Balance int
}

// Add user data
func (u *User) userData() {
	fmt.Printf("Enter name of user: ")
	var name string
	name, err := io()
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Printf("Enter account number of user: ")
	var number string
	number, err = io()
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Printf("Enter balance: ")
	var balance int
	fmt.Scanf("%d", &balance)

	// Fill the User struct with the entered data
	u.Name = name
	u.Number = number
	u.Balance = balance

	// Store User instance in the map
	m[u.Number] = u
}

// Withdraw money
func (u *User) withdraw(amount int) {
	if u.Balance-amount < 0 {
		fmt.Println("Can't be withdrawn")
	} else {
		u.Balance -= amount
		fmt.Printf("Amount Rs. %d is withdrawn\n", amount)
	}
}

// Add money
func (u *User) add(amount int) {
	u.Balance += amount
	fmt.Printf("Amount Rs. %d is added\n", amount)
}

// Display all users
func details() {
	for _, u := range m {
		fmt.Printf("Account name: %s | Account number: %s | Balance: %d\n", u.Name, u.Number, u.Balance)
	}
}

func main() {
	var n int

	for {
		fmt.Printf("1. Add user\n2. Add money\n3. Remove money\n4. View account details\n0. Exit\n")

		fmt.Printf("Enter choice:")
		fmt.Scanf("%d", &n)

		switch n {
		case 1:

			newUser := &User{} // New instance per user
			newUser.userData()

		case 2: // Add money
			fmt.Printf("Enter account number: ")
			number, err := io()
			if err != nil {
				fmt.Println("Error:", err)
				break
			}

			v, exist := m[number]
			if !exist {
				fmt.Println("Error: user not found")
				break
			}

			fmt.Printf("Money to add: ")
			amountStr, _ := io()
			amount, err := strconv.Atoi(amountStr)
			if err != nil {
				fmt.Println("Invalid amount")
				break
			}

			v.add(amount)

		case 3: // Withdraw money
			fmt.Printf("Enter account number: ")
			number, err := io()
			if err != nil {
				fmt.Println("Error:", err)
				break
			}

			v, exist := m[number]
			if !exist {
				fmt.Println("Error: user not found")
				break
			}

			fmt.Printf("Money to withdraw: ")
			amountStr, _ := io()
			amount, err := strconv.Atoi(amountStr)
			if err != nil {
				fmt.Println("Invalid amount")
				break
			}

			v.withdraw(amount)
		case 4: //view details
			details()
		}
		if n == 0 {
			break

		}
	}
}

// Buffered io for input
func io() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	text = strings.TrimSpace(text)
	return text, nil
}
