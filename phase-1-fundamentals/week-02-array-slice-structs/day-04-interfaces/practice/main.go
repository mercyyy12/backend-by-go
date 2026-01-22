package main

import "fmt"

// Vehicle interface
type Vehicle interface {
	Start()
	Stop()
	Info()
}

type Car struct {
	Brand     string
	Model     string
	FuelLevel int
	Distance  int
}

type Bike struct {
	Brand     string
	Type      string
	GearCount int
}

type Truck struct {
	Brand    string
	Capacity string
	Load     int
}

func (c *Car) Start() {
	fmt.Println("car is starting...")
}

func (b *Bike) Start() {
	fmt.Println("bike is starting...")
}

func (t *Truck) Start() {
	fmt.Println("truck is starting...")
}

func (c *Car) Stop() {
	fmt.Println("car is stopping...")
}

func (b *Bike) Stop() {
	fmt.Println("bike is stopping...")
}

func (t *Truck) Stop() {
	fmt.Println("truck is stopping...")
}

func (c *Car) Info() {
	fmt.Printf("Brand = %v\nModel = %v\nFuel Level = %v\n", c.Brand, c.Model, c.FuelLevel)
}

func (b *Bike) Info() {
	fmt.Printf("Brand = %v\nType = %v\nGears = %v\n", b.Brand, b.Type, b.GearCount)
}

func (t *Truck) Info() {
	fmt.Printf("Brand = %v\nCapacity = %v\nLoad = %v\n", t.Brand, t.Capacity, t.Load)
}

func main() {
	// Slice of interface
	vehicleRecord := make([]Vehicle, 0)

	c := &Car{Brand: "Toyota", Model: "Corolla", FuelLevel: 10}
	b := &Bike{Brand: "Honda", Type: "Mountain", GearCount: 21}
	t := &Truck{Brand: "Volvo", Capacity: "10 tons", Load: 5}

	vehicleRecord = append(vehicleRecord, c, b, t)

	for _, v := range vehicleRecord {
		v.Start()
		v.Info()
		v.Stop()
	}

}
