package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var listing = make([]string, 0)
var storage = make(map[string]*Item)
var qty int

// Product in inventory
type Item struct {
	name     string
	price    int
	quantity int
	category string
}

// adds a new item to storage
func (i *Item) addItem() {
	fmt.Printf("Enter item name: ")
	i.name = io()

	fmt.Printf("Enter price: ")
	fmt.Scanf("%d\n", &i.price)
	if i.price < 0 {
		println("error") // price can't be negative
		return
	}

	fmt.Printf("Enter category: ")
	i.category = io()

	fmt.Printf("Enter quantity: ")
	fmt.Scanf("%d\n", &i.quantity)

	fmt.Printf("Item %s added successfully!\n", i.name)
	storage[i.name] = i // store pointer in map

	listing = append(listing, i.name) // add to list for ordered display
}

// To increase quantity of an item
func (i *Item) addStock(qty int) {
	i.quantity += qty
	fmt.Printf("Stock updated! %s now has %d units.\n", i.name, i.quantity)
}

// to decrease quantity if enough stock exists
func (i *Item) reduceStock(qty int) error {
	if i.quantity < qty {
		return fmt.Errorf("not enough stock")
	}
	i.quantity -= qty
	fmt.Printf("Stock reduced! %s now has %d units.\n", i.name, i.quantity)
	return nil
}

// prints all items in listing slice
func listItems() {
	fmt.Println("Inventory list")
	var i int = 1
	for _, v := range listing {
		fmt.Printf("%d. %s\t|\t%s\t|\tRs. %d\t|\tQty: %d|\tValue: %d\n", i, storage[v].name, storage[v].category, storage[v].price, storage[v].quantity, storage[v].quantity*storage[v].price)
		i++
	}
}

func main() {

	var n int
	defer fmt.Println("Exiting Inventory manager....") // will run at the end
	fmt.Println("1. Add item\n2. Add stock\n3. Reduce stock\n4. List items\n5. Total inventory value\n6. Delete item\n0. Exit")

	for {
		fmt.Printf("Enter choice: ")
		fmt.Scanf("%d", &n)

		switch n {
		case 1:
			item := &Item{} // create new item
			item.addItem()  // call addItem function

		case 2:
			fmt.Print("Enter item name: ")
			name := io() // read item name

			item, exists := storage[name] // get item from map
			if !exists {
				fmt.Println("Item not found!") // item doesn't exist
				continue
			}

			fmt.Print("Enter quantity to add: ")
			fmt.Scanf("%d\n", &qty)

			item.addStock(qty) // increase stock

		case 3:
			fmt.Print("Enter item name: ")
			name := io()

			item, exists := storage[name]
			if !exists {
				fmt.Println("Item not found!")
				continue
			}

			fmt.Print("Enter quantity to reduce: ")
			fmt.Scanf("%d\n", &qty)
			err := item.reduceStock(qty)
			if err != nil {
				fmt.Println(err) // not enough stock
			}
		case 4:
			listItems() // print all items

		case 5:
			var totalValue = 0
			fmt.Println("Total inventory value = ")
			for _, value := range listing {
				totalValue += storage[value].price * storage[value].quantity
			}
			fmt.Println("Total inventory value:", totalValue)

		case 6:
			fmt.Printf("Enter item name to delete: ")
			name := io()

			delete(storage, name) // remove from map
			fmt.Printf("Item %s removed from inventory\n", name)

		case 0:
			return

		default:
			fmt.Println("Enter valid choice!") // catch invalid input
		}

		if n == 0 {
			break
		}
	}
}

// TO read input from user and trims spaces/lowercases
func io() string {
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("error", err)
		return "1"
	}
	name = strings.TrimSpace(strings.ToLower(name))
	return name
}
