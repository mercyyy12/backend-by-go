/*
Library Book Management CLI

A simple program to manage books:
- Add books
- View available books
- Borrow books
- Return books

Demonstrates:
- Structs & nested structs
- Slices of pointers
- Pointer receiver methods
- Basic user input handling
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// book structure with nested struct
type books struct {
	id     int
	title  string
	author string
	status // nested structure
}

type status struct {
	isborrowed bool // to check if book is borrowed or not
	borrower   string
}

// Globally declared slice of pointers to book
var bookRecord = make([]*books, 0)
var c = counter()

// A method to add books
func (b *books) addBooks(bookName string, authorName string) {
	b.id = c()
	b.title = bookName
	b.author = authorName
	b.isborrowed = false               // new books are not borrowed
	b.borrower = ""                    // initially no borrow
	bookRecord = append(bookRecord, b) // adding a book pointer in slice
}

// Method to borrow a book
func (b *books) borrowBook(name string) {
	b.borrower = name
	b.isborrowed = true // mark book is borrowed
}

// Method to return a book
func (b *books) returnBook() {
	b.isborrowed = false // mark the book as available
	b.borrower = ""      // clear borrower name
}

func main() {
	user := &books{} // reusable book object

	for {
		fmt.Printf("Enter a number: \n1. To add a book\n2. To view available books\n3. To borrow a book\n4. To return a book\n0. To exit\n")
		n := inputInt()

		switch n {
		case 1: // add a book
			user = &books{} // create a new book object
			fmt.Print("Enter the title of the book: ")
			bookName := input()

			fmt.Print("Enter the author's name: ")
			authorName := input()

			user.addBooks(bookName, authorName)

		case 2: // view available books

			for _, v := range bookRecord {
				if v.isborrowed == false { // only show books that are not borrowed
					fmt.Printf("The id of the book is: %d, title = %v\n", v.id, v.title)
				}

			}

		case 3: // borrow a book
			var bookToBorrow *books
			fmt.Printf("Enter the ID of the book that you want to borrow: ")
			id := inputInt()
			for _, v := range bookRecord {
				if v.id == id {
					bookToBorrow = v
					break
				}
			}

			if bookToBorrow == nil {
				fmt.Println("Book not found!")
				continue
			}

			if bookToBorrow.isborrowed {
				fmt.Printf("%v is already borrowed by %v\n", bookToBorrow.title, bookToBorrow.borrower)
				continue // back to menu
			}

			fmt.Print("Enter your name: ")
			name := input()
			bookToBorrow.borrowBook(name)
			fmt.Printf("%s successfully borrowed book %v\n", bookToBorrow.borrower, bookToBorrow.title)

		case 4:
			fmt.Printf("Enter the ID of the book that you want to return: ")
			id := inputInt()

			for _, v := range bookRecord {
				if v.id == id && v.isborrowed == true {
					{
						fmt.Printf("%v is returned by %v\n", v.title, v.borrower)
						v.returnBook()
					}
				}
			}
		case 0: // exit program
			return
		}

	}
}

// input function to read a line from user
func input() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text1 := strings.ToLower(strings.TrimSpace(text))
	return text1
}

func inputInt() int {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	num, err := strconv.Atoi(text)
	if err != nil {
		fmt.Println("Please enter a integer")
		return -1
	}
	return num
}

func counter() func() int {
	count := 0
	return func() int {
		result := count // store current value
		count++         // increment for next call
		return result   // return current value
	}
}
