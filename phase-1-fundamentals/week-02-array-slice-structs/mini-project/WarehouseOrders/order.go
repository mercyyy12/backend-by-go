package main

import "fmt"

// Generates unique order ID
var nextOrderID = orderIDCounter()

// Stores orders by order ID
var orders = make(map[int]*Order)

// a product with quantity
type OrderItem struct {
	Product  *Product
	Quantity int
}

// Assigns product and quantity to order item
func (o *OrderItem) AssignProduct(product *Product, quantity int) {
	o.Product = product
	o.Quantity = quantity
}

// a customer order
type Order struct {
	TotalPrice int
	OrderID    int
	Items      []*OrderItem
	OrderStatus
}

func (o *Order) totalAmount(itemNum int, cost *Product) {
	o.TotalPrice += cost.Price * itemNum
}

// Adds an item to an order
func (o *Order) AddItem(id int, item *OrderItem) {
	o.OrderID = id
	o.Items = append(o.Items, item)
}

func (o Order) PrintItems() {
	for _, v := range o.Items {
		fmt.Printf("Name: %v\tQuantity: %v\n", v.Product.Name, v.Quantity)
	}
}

// Tracks order state
type OrderStatus struct {
	pending   bool
	delivered bool
}

// Generates order IDs for orders starting from 500
func orderIDCounter() func() int {
	count := 500
	return func() int {
		result := count
		count++
		return result
	}
}
