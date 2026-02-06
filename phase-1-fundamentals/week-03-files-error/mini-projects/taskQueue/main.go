package main

import (
	"fmt"
	"sync"
	"time"
)

// ID generator function for tasks
var c = countId()

// used the 'data' interface to hold any task type
var storedata = make(map[int]data)

// worker goroutine to process task
func goWorkers(id int, ch <-chan data) {

	for j := range ch {
		i, v := j.worker()
		fmt.Printf("worker %v picked up task %v (%v) \n", id, i, v)
		time.Sleep(2 * time.Second) // simulating processing time
	}
	fmt.Printf("worker %v completed task successfully!\n", id)
}

func main() {

	fmt.Println("Produce Consumer Task Queue")

loop:
	for {
		fmt.Printf("Options:\n1. Add task\n2. List tasks\n3. Start processing\n4. Exit\nEnter a number: ")
		menuOptions, _ := inputInt()
		switch menuOptions {
		case 1:
			for {
				fmt.Printf("Enter task type (email / image / report): ")
				taskType, _ := input()

				switch taskType {
				case "email":
					task := &structEmail{}
					fmt.Printf("Enter payload (to:msg): ")
					payload, _ := input()
					id := task.addEmail(payload)
					fmt.Printf("Task added (ID: %d)\n", id)

				case "image":
					task := &structImage{}
					fmt.Printf("Enter payload (file:resize): ")
					payload, _ := input()
					id := task.addImage(payload)
					fmt.Printf("Task added (ID: %d)\n", id)

				case "report":
					task := &structReport{}
					fmt.Printf("Enter payload (title:content): ")
					payload, _ := input()
					id := task.addReport(payload)
					fmt.Printf("Task added (ID: %d)\n", id)

				default:
					fmt.Println("Invalid task type, try again.")
					continue
				}

				for {
					fmt.Printf("Would you like to add another task? (y/n): ")
					choice, _ := input()
					if choice == "y" {
						break
					} else if choice == "n" {
						goto loop
					} else {
						fmt.Println("Invalid choice. Enter y or n.")
					}
				}
			}

		case 2:
			// list all task
			for _, v := range storedata {
				fmt.Println(v.showData())
			}

		case 3:
			// start processing tasks withs workers
			start := time.Now()
			var wg sync.WaitGroup
			var ch = make(chan data, 10)
			fmt.Println("launching 5 workers......")

			// start 5 worker goroutines
			for w := 1; w <= 5; w++ {
				id := w
				wg.Go(func() {
					goWorkers(id, ch)
				})
			}

			// send task to channels
			wg.Go(func() {
				for _, v := range storedata {
					ch <- v
				}
				close(ch)
			})

			wg.Wait() // waits for all workers to finish
			fmt.Printf("All tasks finished!\n Total runtime: %v\n", time.Since(start).Seconds())

		case 4:
			fmt.Println("Exiting program...")
			return // exit main

		default:
			fmt.Println("Invalid option. Enter 1–4.")
		}

	}
}
