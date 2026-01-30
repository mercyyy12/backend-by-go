package main

import (
	"fmt"
	"sync"
)

func main() {

	// declare waitgroup
	var wg sync.WaitGroup

	// waiting for 1 Goroutine
	wg.Add(1) // outside goroutine
	go func() {
		for range 5 {
			fmt.Println("1st example")
		}
		defer wg.Done() // mark goroutine as done (should be inside goroutine)
	}()

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Println(n)
		}(i)
	}

	// Starts a goroutine and manages Add() and Done() auto
	wg.Go(func() {
		println("2nd example")
	})

	fmt.Println("from main")

	wg.Wait() // block the main untill all wg.Done() is called
}
