package main

// Generates unique product ID
var nextProductID = productIDCounter()

// Stores products by product ID
var products = make(map[int]*Product)

// represents an item in warehouse
type Product struct {
	ID            int
	Name          string
	Category      string
	Price         int
	StockQuantity int
}

// Adds a product to inventory
func (p *Product) AddToInventory(price, stockQuantity int, name, category string) {
	p.ID = nextProductID()
	p.Price = price
	p.StockQuantity = stockQuantity
	p.Name = name
	p.Category = category
	products[p.ID] = p
}

// Reduces product stock after ordering
func (p *Product) ReduceStock(value int) {
	p.StockQuantity -= value
}

// Generates incremental IDs for products starting from 0
func productIDCounter() func() int {
	count := 1
	return func() int {
		result := count
		count++
		return result
	}
}
