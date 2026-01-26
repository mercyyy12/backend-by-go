package main

// Generates unique customer ID
var nextCustomerID = customerIDCounter()

// Stores customers by customer ID
var customers = make(map[int]*Customer)

// Customer
type Customer struct {
	ID     int
	Name   string
	Email  string
	Orders []*Order
}

// Registers a new customer
func (c *Customer) RegisterCustomer(name, email string) int {
	c.Name = name
	c.ID = nextCustomerID()
	c.Email = email
	customers[c.ID] = c
	return c.ID
}

// Adds an order to customer history
func (c *Customer) AddOrder(order *Order) {
	c.Orders = append(c.Orders, order)
}

// Generates incremental IDs for customers starting from 0
func customerIDCounter() func() int {
	count := 1
	return func() int {
		result := count
		count++
		return result
	}
}
