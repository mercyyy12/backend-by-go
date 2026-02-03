package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var c = countId()
var ch = make(chan int)

var storeEmail = make(map[int]*structEmail)
var storeImage = make(map[int]*structImage)
var storeReport = make(map[int]*structReport)
var storeId = make([]int, 0)

type structEmail struct {
	emailId   int
	emailTime time.Time
	Payload   string
}

type structImage struct {
	imageId   int
	imageTime time.Time
	Payload   string
}

type structReport struct {
	reportId   int
	reportTime time.Time
	Payload    string
}

func (e *structEmail) addEmail(payload string) int {
	e.emailId = c()
	e.emailTime = time.Now()
	e.Payload = payload
	storeEmail[e.emailId] = e
	storeId = append(storeId, e.emailId)
	return e.emailId
}

func (i *structImage) addImage(payload string) int {
	i.Payload = payload
	i.imageTime = time.Now()
	i.imageId = c()
	storeImage[i.imageId] = i
	storeId = append(storeId, i.imageId)
	return i.imageId
}

func (r *structReport) addReport(payload string) int {
	r.Payload = payload
	r.reportTime = time.Now()
	r.reportId = c()
	storeReport[r.reportId] = r
	storeId = append(storeId, r.reportId)
	return r.reportId
}

func goWorkers() {

	for i := 1; i <= 5; i++ {
		// checks if
		go func() {
			val, ok := <-ch // receive from channel
			if !ok {        // ok == false = channel is closed
				fmt.Println("channel has closed")
				// break
			}
			fmt.Printf("worker %d picked up task %d \n", i, val)
		}()
	}
}

func main() {
	fmt.Println("Produce Consumer Task Queue")
	for {
		fmt.Printf("Enter:\n1. Add task\n2. List tasks\n3. Start processing\n4. Exit\nEnter a number: ")
		menuOptions, _ := inputInt()
		switch menuOptions {
		case 1:
			fmt.Printf("Enter task type (email / image / report): ")
			taskType, _ := input()

			switch taskType {
			case "email":
				task := &structEmail{}
				fmt.Printf("Enter payload (to:msg): ")
				payload, _ := input()
				id := task.addEmail(payload)
				fmt.Printf("Task added (ID: %d)", id)

			case "image":
				task := &structImage{}
				fmt.Printf("Enter payload (file:resize): ")
				payload, _ := input()
				id := task.addImage(payload)
				fmt.Printf("Task added (ID: %d)", id)

			case "report":
				task := &structReport{}
				fmt.Printf("Enter payload (title:content): ")
				payload, _ := input()
				id := task.addReport(payload)
				fmt.Printf("Task added (ID: %d)", id)
			}

		case 2:
			for _, v := range storeId {
				val1, exist := storeImage[v]
				if !exist {
					val, exist := storeEmail[v]
					if !exist {
						val2, exist := storeReport[v]
						if !exist {
							continue
						} else {
							fmt.Printf("ID: %d | Type: | Time: %v | Payload: %s\n", val2.reportId, val2.reportTime.Format("15:04:05"), val2.Payload)
						}
					} else {
						fmt.Printf("ID: %d\t|\tType: \t|\tTime: %v\t|\tPayload: %s\n", val.emailId, val.emailTime.Format("15:04:05"), val.Payload)
					}
				} else {
					fmt.Printf("ID: %d\t|\tType: \t|\tTime: %v\t|\tPayload: %s\n", val1.imageId, val1.imageTime.Format("15:04:05"), val1.Payload)
				}
			}

		case 3:
			goWorkers()

			for i, v := range storeId {

				fmt.Println("sending data", i)
				ch <- v
			}
			// closing a channel after sending all values
			close(ch)

		}

	}
}

func input() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return "thik xaina", fmt.Errorf("Input lida erro %w", err)
	}
	text = strings.TrimSpace(text)
	return text, nil
}

func inputInt() (int, error) {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return -1, fmt.Errorf("integer Input lida error %w", err)
	}

	text = strings.TrimSpace(text)
	num, err := strconv.Atoi(text)
	if err != nil {
		return 0, fmt.Errorf("int conversion ma error")
	}
	return num, nil
}

func countId() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}
