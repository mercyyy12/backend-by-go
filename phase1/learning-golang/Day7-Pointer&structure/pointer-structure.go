package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// struct definitions
type book struct {
	title  string
	author string
	price  int
}

type car struct {
	model string
	price int
}

type location struct {
	city    string
	country string
}

type user struct {
	name string
	age  int
	location
}

type student struct {
	name string
}

type book1 struct {
	title  []string
	author []string
	price  []int
}

// helper func: user info using struct embedding
func userInfo() {
	var p user
	p.name = "ram"
	p.city = "kathmandu"
	p.age = 20
	p.country = "Nepal"
	fmt.Printf("The name of the user is %s, age is %d, and lives in %s, %s\n", p.name, p.age, p.city, p.country)
}

// helper func: input car details and change attributes using pointer
func carip() {
	var p car
	fmt.Println("Enter model and price of the car")
	fmt.Scanf("%s%d", &p.model, &p.price)
	fmt.Printf("Before: model=%s price=%d\n", p.model, p.price)
	carAttChange(&p)
	fmt.Printf("After:  model=%s price=%d\n", p.model, p.price)
}

// helper func: mutate car fields via pointer
func carAttChange(c *car) {
	c.model = "pulsar ns 200"
	c.price = 456000
}

// helper func: input slice of books
func slicee() {
	var p book1
	var num int
	fmt.Print("How many book records? ")
	fmt.Scanf("%d", &num)

	p.title = make([]string, num)
	p.author = make([]string, num)
	p.price = make([]int, num)

	for i := 0; i < num; i++ {
		fmt.Println("Enter title author price:")
		fmt.Scanf("%s%s%d", &p.title[i], &p.author[i], &p.price[i])
	}

	for i := range p.title {
		fmt.Printf("Book %d: %s by %s costs %d\n", i+1, p.title[i], p.author[i], p.price[i])
	}
}

// helper func: swap using pointers
func swap(x, y *int) {
	*x, *y = *y, *x
}

// helper func: double a number using pointer
func double(x *int) {
	*x *= 2
}

// helper func: demonstrate value vs pointer receiver with struct
func ipstruct() {
	var p student
	fmt.Print("Enter name: ")
	fmt.Scan(&p.name)

	fmt.Printf("Original name: %s\n", p.name)
	valueSwap(p)
	fmt.Printf("After valueSwap: %s\n", p.name)
	pointerSwap(&p)
	fmt.Printf("After pointerSwap: %s\n", p.name)
}

// helper func: struct passed by value
func valueSwap(s student) {
	s.name = "ram2"
}

// helper func: struct passed by pointer
func pointerSwap(s *student) {
	s.name = "ram1"
}

// helper func: input two integers
func ip() (int, int) {
	var x, y int
	fmt.Print("Enter two numbers: ")
	fmt.Scanf("%d%d", &x, &y)
	return x, y
}

// helper func: input single book using bufio and string split
func ipBook() {
	var p book
	fmt.Println("Enter title, author and price of the book")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fields := strings.Fields(strings.TrimSpace(input))
	if len(fields) != 3 {
		fmt.Println("Expected 3 values: title author price")
		return
	}

	p.title = fields[0]
	p.author = fields[1]
	p.price, err = strconv.Atoi(fields[2])
	if err != nil {
		fmt.Println("Invalid price:", err)
		return
	}

	fmt.Printf("Book: %s by %s costs %d\n", p.title, p.author, p.price)
}

func main() {
	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1: Swap two ints")
		fmt.Println("2: Pointer to variable (double)")
		fmt.Println("3: Struct creation & printing (book)")
		fmt.Println("4: Pass struct by value vs pointer")
		fmt.Println("5: Slice of structs")
		fmt.Println("6: Pointer to struct (car)")
		fmt.Println("7: Struct embedding (user/location)")
		fmt.Println("0: Exit")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			a, b := ip()
			fmt.Printf("Before swap: a=%d b=%d\n", a, b)
			swap(&a, &b)
			fmt.Printf("After  swap: a=%d b=%d\n", a, b)
		case 2:
			num := 10
			fmt.Printf("Before: %d\n", num)
			double(&num)
			fmt.Printf("After:  %d\n", num)
		case 3:
			ipBook()
		case 4:
			ipstruct()
		case 5:
			slicee()
		case 6:
			carip()
		case 7:
			userInfo()
		case 0:
			fmt.Println("Good-bye!")
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}
