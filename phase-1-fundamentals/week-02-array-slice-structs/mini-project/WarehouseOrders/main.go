package main

import (
	"fmt"
)

func main() {
	// Infinite loop to keep program running
	for {
		fmt.Printf("Select:\n1. Add a product\n2. Add a customer\n3. List all product\n4. List all customer\n5. Place an order\nEnter: ")
		n := inputInt()

		switch n {

		// Add products to warehouse
		case 1:
			addProducts()

		// Register a new customer
		case 2:
			registerCustomer()

		// List all products
		case 3:
			fmt.Println("---Listing available products---")
			for _, v := range products {
				fmt.Printf("ID: %v\tName: %v\tCategory: %v\tPrice: %v\tStock: %v\n",
					v.ID, v.Name, v.Category, v.Price, v.StockQuantity)
			}

		// List all customers
		case 4:
			fmt.Println("---Listing customers---")
			for _, v := range customers {
				fmt.Printf("ID: %v\tName: %v\n", v.ID, v.Name)
			}

		// Place an order
		case 5:
			createOrder()
		case 6:
			return
		default:
			fmt.Println("Enter valid number")
		}
	}
}

func addProducts() {
	fmt.Println("---Adding products to warehouse---")
	for {
		product := &Product{}

		fmt.Printf("Name: ")
		name := input()

		fmt.Printf("Category: ")
		category := input()

		fmt.Printf("Price: ")
		price := inputInt()

		fmt.Printf("Stock: ")
		stock := inputInt()

		product.AddToInventory(price, stock, name, category)

		fmt.Printf("Would you like to add another product?\nEnter 'yes' or 'no': ")
		validate := input()
		if validate != "yes" {
			break
		}
	}
}

func registerCustomer() {
	fmt.Println("---Registering customer---")
	customer := &Customer{}

	fmt.Printf("Name: ")
	name := input()

	fmt.Printf("Email: ")
	email := input()

	// Generates unique ID and stores customer
	id := customer.RegisterCustomer(name, email)
	fmt.Println("Your ID is: ", id)
}

func createOrder() {
	fmt.Println("---Creating a new order---")

	fmt.Printf("Enter your user id: ")
	id := inputInt()

	// Validates if customer exists
	customer, exist := customers[id]
	if !exist {
		fmt.Println("the user doesn't exist")
		return
	}

	// Generate order ID
	orderID := nextOrderID()
	fmt.Printf("Order created: OrderId %v for customer %v\n", orderID, customer.Name)

	order := &Order{}
	order.pending = true

	getOrderItem(order, orderID)

	orders[orderID] = order
	customer.AddOrder(order)

	fmt.Printf("Order ID: %v\n", orderID)
	order.PrintItems()
	fmt.Printf("Total price = %v\n", order.TotalPrice)

	chooseDelivery(order)
}

func getOrderItem(order *Order, orderID int) {
	// Add multiple products to the same order
	for {
		orderItem := &OrderItem{}

		fmt.Printf("Enter the id of the product that you would like to purchase: ")
		id := inputInt()

		// Validate product existence
		product, exists := products[id]
		if !exists {
			fmt.Println("the product doesn't exist")
			continue
		}

		fmt.Printf("Enter the number of items that you want to buy: ")
		itemNum := inputInt()

		if product.StockQuantity < itemNum {
			fmt.Println("The order amount exceeds the available amount")
			continue
		}

		// Reduce stock and add to order
		product.ReduceStock(itemNum)
		order.totalAmount(itemNum, product)

		orderItem.AssignProduct(product, itemNum)
		order.AddItem(orderID, orderItem)

		fmt.Printf("Would you like to add another product?\nEnter 'yes' or 'no': ")
		validate := input()
		if validate == "no" {
			break
		}
	}
}

func chooseDelivery(order *Order) {
	// Select delivery type
	fmt.Println("---Choose the type of delivery--")
	fmt.Println("Enter 1 for basic and 2 for premium")
	n := inputInt()
	var d Delivery
	switch n {
	case 1:
		d = BasicDelivery{}
		d.Deliver(order)

	case 2:
		d = PremiumDelivery{}
		d.Deliver(order)
	}
}
