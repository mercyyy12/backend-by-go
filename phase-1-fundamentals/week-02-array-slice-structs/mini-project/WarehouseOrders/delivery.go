package main

import "fmt"

// Deliver interface to handle a delivery
type Delivery interface {
	Deliver(order *Order)
}

// Basic delivery type
type BasicDelivery struct {
}

// Premium delivery type
type PremiumDelivery struct {
}

// Basic delivery
func (b BasicDelivery) Deliver(order *Order) {
	order.delivered = true
	order.pending = false
	fmt.Printf("Order ID: %v is set to be sent by basic delivery\n", order.OrderID)
}

// Premium delivery
func (p PremiumDelivery) Deliver(order *Order) {
	order.delivered = true
	order.pending = false
	fmt.Printf("Order ID: %v is set to be sent by premium delivery\n", order.OrderID)
}
