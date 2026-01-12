package main

import (
	"fmt"
	"time"
)

func main() {

	// creating buffered channel
	ch := make(chan int, 5)           // capacity 5
	fmt.Println("length:", len(ch))   // items in buffer
	fmt.Println("capacity:", cap(ch)) // total capacity

	// send to buffered channel
	fmt.Println("sending")
	ch <- 25
	ch <- 35
	ch <- 45

	fmt.Println("length:", len(ch)) // items in buffer

	// range over ch (receiving)
	// for i := range ch {
	// 	fmt.Println(i)
	// }

	go func() {
		ch <- 25
		ch <- 35
		ch <- 45
		ch <- 55
		ch <- 65
		ch <- 75
		close(ch) // close buffered channel
	}()

	// receive using ok
	for {
		val, ok := <-ch
		if !ok {
			fmt.Println("Channel closed")
			break
		} else {
			fmt.Println(val)
		}
	}

	ch1 := make(chan int, 2)
	i := 1

	// producer-consumer pattern
	ch2 := make(chan int, 3)

	//producer
	go func() {
		for j := range 3 {
			ch2 <- j
		}
		close(ch2)
	}()

	// consumer
	for j := range ch2 {
		fmt.Println(j)
	}

	// Non blocking sender
	go func() {
		for i <= 10 {
			select {
			case ch1 <- i:
				fmt.Println("sent:", i)
				i++
			default:
				fmt.Println("Buffer full, waiting")
				time.Sleep(100 * time.Millisecond)
			}
		}
		close(ch1)
	}()

	// Non blocking receiver
	for {
		select {
		case val, ok := <-ch1:
			if !ok {
				fmt.Println("Channel closed")
				return
			}
			fmt.Println("received", val)
		default:
			fmt.Println("Buffer empty, waiting")
			time.Sleep(100 * time.Millisecond)
		}
	}

}
