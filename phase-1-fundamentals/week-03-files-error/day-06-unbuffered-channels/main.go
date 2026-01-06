package main

import (
	"fmt"
	"time"
)

func main() {

	// creating unbuffered channel
	ch := make(chan int)
	ch1 := make(chan int)
	fmt.Printf("Channel created: %v\n", ch) // Memory address of ch

	// sending a single value into channel
	go func() {
		fmt.Println("sending data.....")
		ch <- 45 // sending data (blocks untill main receives)
		fmt.Println("data sent!")
	}()

	// receives the data from ch
	fmt.Println("Receiving data....")
	// msg := <-ch
	fmt.Printf("The msg received is %v\n", <-ch) // go routine is unblocked (handshaking)

	// sending multiple data into a channel
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("sending data")
			ch1 <- i
		}

		// closing a channel after sending all values
		close(ch1)
	}()

	// Range over ch1
	// for i := range ch1 {
	// 	fmt.Println("Data sent:", i)
	// }

	// receive with ok
	for {
		// checks if
		val, ok := <-ch1 // receive from channel
		if !ok {         // ok == false = channel is closed
			fmt.Println("channel has closed")
			break
		}
		fmt.Println("Data sent: ", val)
	}

	// wait a bit so goroutines can finish
	time.Sleep(time.Second)
}
