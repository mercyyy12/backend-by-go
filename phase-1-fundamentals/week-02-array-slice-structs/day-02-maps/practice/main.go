/*
User Management CLI in Go

A simple command line program to manage users with the following features:
1. Add new users (ID, name, email)
2. Update user details (name or email)
3. Delete users by ID
4. Print all users
5. Exit

Topics/Concepts used:
- Structs and pointer receivers
- Maps with pointers to structs
- Methods and function
- Input handling with bufio and fmt
- Time package for timestamps
- Loops, switch statements, and error handling
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Globally declared map which points to user struct
var data = make(map[int]*user)

type user struct {
	id      int
	name    string
	email   string
	created time.Time
}

// prompts to add user to the map
func (u *user) addUser() {
	// var id int
	var err error

	// loop untill valid integer
	for {
		fmt.Print("Enter id: ")
		num, _ := input()
		u.id, err = strconv.Atoi(num)
		if err != nil {
			fmt.Println("It is not a number!", u.id, err)
		} else {
			break
		}
	}

	fmt.Printf("Enter name: ")
	name, err := input()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Printf("Enter email: ")
	email, err := input()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	created := time.Now()

	// assign values to struct and store in map
	u.name = name
	u.created = created
	u.email = email
	data[u.id] = u

}

// to modify user name or email
func (u *user) update() {
	var choose int

	fmt.Println("Enter 1 to change name\nEnter 2 to change email")
	fmt.Scanf("%d", &choose)
	switch choose {
	case 1:
		fmt.Printf("Enter new name: ")
		name, err := input()
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		u.name = name // updates name

	case 2:
		fmt.Printf("Enter email: ")
		email, err := input()
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		u.email = email //updates email
	}
}

// To delete user from the map
func (u *user) deleteData() {
	delete(data, u.id)
}

func main() {

	users := &user{} //pointer to hold new user structs

	for {
		fmt.Printf("Enter:\n1. To add a new user\n2. Update details of a user\n3. To delete a user\n4.Print all user\n5. Exit\nEnter a number: ")

		var number int
		var err error

		// loop untill valid number is entered
		for {
			num, _ := input()
			number, err = strconv.Atoi(num)
			if err != nil && number < 1 && number > 5 {
				fmt.Println("It is not a number!", number, err)
			} else {
				break
			}
			fmt.Println("Invalid input, please enter a number from 1 to 5")
		}

		switch number {
		case 1:

			users = &user{} // create a new user
			users.addUser() // add user

		case 2:
			var update int
			fmt.Print("Enter the id of the user that you want to update: ")
			fmt.Scanf("%d", &update)

			// check if the user exist or not in map
			u, exists := data[update]
			if !exists {
				fmt.Printf("No user with ID = %d", update)
				return
			}
			u.update() // update user

		case 3:
			var deleteId int
			fmt.Print("Enter the id of the user that you want to delete: ")
			fmt.Scanf("%d", &deleteId)

			// checks if the user exist or not in map
			u, exists := data[deleteId]
			if !exists {
				fmt.Printf("No user with ID = %d", deleteId)
				return
			}
			u.deleteData()

		case 4:
			// Print all user
			for _, v := range data {
				fmt.Printf("ID = %v\tName = %v\tEmail = %v\tCreated time = %v\n", v.id, v.name, v.email, v.created)
			}

		case 5:
			// program exit
			return
		}
	}

}

// to take user input
func input() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("Error while taking input: %w", err)
	}
	text = strings.ToLower(strings.TrimSpace(text))
	return text, nil
}
