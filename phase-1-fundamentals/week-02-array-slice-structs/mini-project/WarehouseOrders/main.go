package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Deliver interface {
	Delivery(order *Order) string
}

type Basic struct {
}

type premium struct {
}

func (b Basic) Delivery(order *Order) string {
	order.delivered = true
	order.pending = false
	return fmt.Sprintf("Order ID: %v is set to be sent by basic delivery", order.OrderId)
}

func (p premium) Delivery(order *Order) string {
	order.delivered = true
	order.pending = false
	return fmt.Sprintf("Order ID: %v is set to be sent by premium delivery", order.OrderId)
}

var orderCount = orderIdCounter()
var count = counter()
var productMap = make(map[int]*Product)
var customerMap = make(map[int]*Customer)
var ordertMap = make(map[int]*Order)

type Product struct {
	ID            int
	Name          string
	Category      string
	Price         int
	StockQuantity int
}

func (p *Product) addProduct(Id, Price, StockQuantity int, Name, Category string) {
	p.ID = Id
	p.Price = Price
	p.StockQuantity = StockQuantity
	p.Name = Name
	p.Category = Category
	productMap[Id] = p
}

type Customer struct {
	Id     int
	Name   string
	Email  string
	Orders []*Order
}

func (c *Customer) addCustomer(Name, Email string) int {
	c.Name = Name
	c.Id = count()
	c.Email = Email
	customerMap[c.Id] = c
	return c.Id
}

func (c *Customer) orderedItem(value *Order) {
	c.Orders = append(c.Orders, value)
}

// func (c Customer) showOrder(quantity *Order) {
// 	for _, v := range c.Orders {
// 		fmt.Printf("cName: %v\tQuantity: %v\n", v.OrderId, quantity.Quantity)
// 	}
// }

type Order struct {
	OrderId int
	Items   []*OrderItem
	status
}

type OrderItem struct {
	Product  *Product
	Quantity int
}

func (o *Order) ordered(id int, value *OrderItem) {
	o.OrderId = id
	o.Items = append(o.Items, value)
}

func (o *OrderItem) itemsOrder(value *Product, quantity int) {
	o.Product = value
	o.Quantity = quantity
}

// func (o Order) showOrderProducts(quantity *Order) {
// 	for _, v := range o.Products {
// 		fmt.Printf("oName: %v\tQuantity: %v\n", v.Name, quantity)
// 	}
// }

func (p *Product) stockModify(value int) {
	p.StockQuantity -= value
}

type status struct {
	pending bool
	// shipped   string
	delivered bool
}

func main() {
	// user := make([]*Product, 0)

	// var n int
	for {
		fmt.Printf("Select:\n1. Add a product\n2. Add a customer\n3. List all product\n4. List all customer\n5. Place an order\nEnter: ")
		n := inputInt()

		switch n {
		case 1:

			fmt.Println("---Adding products to warehouse---")
			for {
				users := &Product{}
				fmt.Printf("ID: ")
				id := inputInt()

				fmt.Printf("Name: ")
				name := input()

				fmt.Printf("Category: ")
				category := input()

				fmt.Printf("Price: ")
				price := inputInt()

				fmt.Printf("Stock: ")
				stock := inputInt()

				users.addProduct(id, price, stock, name, category)

				fmt.Printf("Would you like to add another product?\nEnter 'yes' or 'no': ")
				validate := input()
				if validate != "yes" {
					break
				}
			}

		case 2:
			fmt.Println("---Registering customer---")
			users := &Customer{}

			fmt.Printf("Name: ")
			name := input()

			fmt.Printf("Email: ")
			email := input()

			id := users.addCustomer(name, email)
			fmt.Println("Your ID is: ", id)

		case 3:
			fmt.Println("---Listing available products---")
			for _, v := range productMap {
				fmt.Printf("ID: %v\tName: %v\tCategory: %v\tPrice: %v\tStock: %v\n", v.ID, v.Name, v.Category, v.Price, v.StockQuantity)
			}

		case 4:
			fmt.Println("---Listing customers---")
			for _, v := range customerMap {
				fmt.Printf("ID: %v\tName: %v\n", v.Id, v.Name)
			}

		case 5:

			fmt.Println("---creating a new order---")

			fmt.Printf("Enter your user id: ")
			id := inputInt()

			k, exist := customerMap[id]
			if !exist {
				fmt.Println("the user doesn't exist")
				continue
			}

			orderNum := orderCount()
			fmt.Printf("Order created: OrderId %v for customer %v\n", orderNum, k.Name)

			// var itemNum int
			var value *Product
			var totalAmount = 0
			order := &Order{}
			order = &Order{}
			order.pending = true
			order.delivered = false
			orderitem := &OrderItem{}

			for {
				orderitem = &OrderItem{}
				fmt.Printf("Enter the id of the product that you would like to purchase: ")
				id = inputInt()
				var exists bool
				value, exists = productMap[id]
				if !exists {
					fmt.Println("the product doesn't exist")
					continue
				}

				fmt.Printf("Enter the number of items that you want to buy: ")
				itemNum := inputInt()

				if value.StockQuantity < itemNum {
					fmt.Println("The order amount exceeds the available amount")
					continue
				}
				fmt.Println("\nAdding products to order...")
				value.stockModify(itemNum)

				totalAmount += itemNum * value.Price
				orderitem.itemsOrder(value, itemNum)
				// fmt.Printf("Order ID: %v\nProduct ID: %v\nRequested quantity: %v\nTotal price = %v\n", orderNum, id, itemNum, itemNum*value.Price)
				order.ordered(orderNum, orderitem)
				k.orderedItem(order)
				// value.stockModify(itemNum)

				fmt.Printf("Would you like to add another product?\nEnter 'yes' or 'no': ")
				validate := input()
				if validate == "no" {
					break
				}
			}

			fmt.Printf("Order ID: %v\nProducts:\n", orderNum)
			// k.showOrder(order)
			// order.showOrderProducts(order)
			fmt.Printf("Total price = %v\n", totalAmount)
			// order.ordered(orderNum, itemNum, value)

			fmt.Println("---Choose the type of delivery--")
			fmt.Println("Enter 1 for basic and 2 for premium")
			n := inputInt()
			var d Deliver
			switch n {
			case 1:
				d = Basic{}
				fmt.Println(d.Delivery(order))
			case 2:
				d = premium{}
				fmt.Println(d.Delivery(order))
			}

		case 6:
			fmt.Println("---Listing all the orders---")
			for i, v := range ordertMap {
				fmt.Println(i, v)
			}
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

func orderIdCounter() func() int {
	count := 500
	return func() int {
		result := count // store current value
		count++         // increment for next call
		return result   // return current value
	}
}
