package main

import (
	"fmt"
	"time"
)

func main() {

	// creating a channel
	ch := make(chan int)
	fmt.Printf("Channel created: %v\n", ch) // Memory address

	go func() {
		fmt.Println("sending data.....")
		ch <- 45
		fmt.Println("data sent!")
	}()

	fmt.Println("Receiving data...")
	msg := <-ch
	fmt.Printf("The msg received is %v\n", msg)

	time.Sleep(time.Second)
}
