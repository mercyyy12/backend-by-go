package main

import (
	"fmt"
	"sync"
)

func main() {

	var mu1 sync.Mutex    // declare Mutex
	var wg sync.WaitGroup // declare WaitGroup

	// shared variable
	balance := 0

	// creates a go routine five times
	for range 5 {

		// starting a goroutine
		wg.Go(func() {
			// loops 10 times
			for i := 0; i < 1000; i++ {
				mu1.Lock() // acquired lock

				// critical section
				balance += 10

				mu1.Unlock() // released lock
			}
		})
	}

	// Another way to do the above process
	// for range 5 {
	// starting a goroutine
	// 	go func() {
	// 		for i := 0; i < 1000; i++ {
	// 			mu1.Lock()
	// 			balance += 10
	// 			mu1.Unlock()
	// 		}
	// 		wg.Done()
	// 	}()
	// }

	wg.Wait()            // block the main untill all wg.Done() is called
	fmt.Println(balance) // 50000
}
