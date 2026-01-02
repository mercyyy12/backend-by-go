package main

import (
	"fmt"
	"time"
)

func main() {
	go greetings()
	fmt.Println("I am from main function")

	// Anon function
	go func() {
		fmt.Println("Goroutine with anon function called!")
	}()

	// Output order may differ each time
	go printNumbers("Goroutine1")
	go printNumbers("Goroutine2")

	fmt.Println("Main function done")
	counter := 1

	// race condition in goRoutine
	for i := 0; i < 500; i++ {
		// anon function using GoRoutine
		go func() {
			counter++
		}()
	}

	// go func() {
	// 	counter++
	// }()
	time.Sleep(1 * time.Second) // wait to let goroutine finish
	fmt.Println("the count is ", counter)
}

func greetings() {
	fmt.Println("Function called using Goroutine!")
}

// Example to show concurrent task
func printNumbers(name string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(name, ":", i)
		time.Sleep(100 * time.Millisecond) // wait to let goroutine finish
	}
}
