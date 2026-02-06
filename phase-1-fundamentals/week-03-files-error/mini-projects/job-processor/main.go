package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"
)

var data = make(map[int]*job)
var c = counter()

type job struct {
	id        int
	timestamp time.Time
	jobType   string
	email
	image
	report
}

type email struct {
	emailAddress string
	message      string
}

type image struct {
	imageName string
	resize    string
}

type report struct {
	reportId string
}

func (j *job) ntg(jobtipe string) {
	j.id = c()
	j.jobType = jobtipe
	j.timestamp = time.Now()

	switch jobtipe {
	case "email":
		e := &email{}
		j.emailAddress, j.message = addData(e)
	case "image":
		i := &image{}
		j.imageName, j.resize = imageData(i)
	case "report":
		r := &report{}
		j.reportId = reportData(r)
	default:
		fmt.Println("Error xa jobtipe func maa")
		return
	}

	data[j.id] = j

}

var ch = make(chan *job, 5)

func addData(e *email) (string, string) {

	fmt.Printf("Enter payload keys (e.g., to=alice@gmail.com message=Hi):\nTo = ")
	emailadd, err := input()
	if err != nil {
		fmt.Println("Error in email name:", err)
		return "", ""
	}

	fmt.Printf("message = ")
	msg, err := input()
	if err != nil {
		fmt.Println("Error in message:", err)
		return "", ""
	}

	e.emailAddress = emailadd
	e.message = msg
	return e.emailAddress, e.message

}

func imageData(i *image) (string, string) {
	fmt.Printf("enter payload keys (e.g., file=photo.jpg resize=500x500):\nfile = ")
	file, err := input()
	if err != nil {
		fmt.Println("Error in image name:", err)
		return "", ""
	}
	fmt.Printf("resize = ")
	resize, err := input()
	if err != nil {
		fmt.Println("Error in resize:", err)
		return "", ""
	}

	i.imageName = file
	i.resize = resize

	return i.imageName, i.resize

}

func reportData(r *report) string {
	fmt.Printf("Enter payload keys:\nReport ID = ")
	id, err := input()
	if err != nil {
		fmt.Println("Error in report:", err)
		return ""
	}
	r.reportId = id
	return r.reportId

}

var panicCounter = 0
var successCounter = 0
var failedCounter = 0

func main() {
	var wg sync.WaitGroup
	start := time.Now()

	defer func() {
		fmt.Println("Runtime:", time.Since(start))
	}()

	var n int
	fmt.Println("--Job Processor--")
	for {
		fmt.Printf("1. Add job\n2. List queued jobs\n3. Start processing\n4. Shutdown\nEnter number: ")
		// time.Sleep(400 * time.Millisecond)

		fmt.Scanf("%d", &n)

		switch n {
		case 1:
			fmt.Printf("Enter job type (email/image/report): ")
			kho := &job{}
			jobType, err := input()
			if err != nil {
				fmt.Println("Error in job type:", err)
				return
			}
			kho.ntg(jobType)

		case 2:
			fmt.Println("Queued jobs: ")
			for i, v := range data {
				fmt.Printf("ID %d - %v - Created: %v\n", i, v.jobType, v.timestamp)
				switch v.jobType {
				case "email":
					fmt.Printf("Payload: To = %v, message = %v\n", v.emailAddress, v.message)
				case "image":
					fmt.Printf("Payload: File = %v, resize = %v\n", v.imageName, v.resize)
				case "report":
					fmt.Printf("Payload: Report ID = %v\n", v.reportId)
				default:
					fmt.Println("error aayo hauu")
					return
				}
			}

		case 3:

			fmt.Println("starting 5 workers...")
			i := 1

			func() {
				for _, j := range data {
					fmt.Printf("Worker %v processing job %v (%v)\n", i, j.id, j.jobType)
					ch <- j
					i++

				}
				close(ch)
				// time.Sleep(5 * time.Second)

			}()

			for i := 1; i <= 5; i++ {
				wg.Add(1)
				go Worker(i)
				wg.Done()
			}

			fmt.Printf("All jobs processed: %d success, %d failed, %d recovered from panic", successCounter, panicCounter, failedCounter)
			wg.Wait()
		case 4:
			fmt.Println("shutting down...")
			return

		}

	}
}

func input() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return "no text", fmt.Errorf("Error taking input: %w", err)

	}
	text = strings.ToLower(strings.TrimSpace(text))
	return text, nil
}

func Worker(n int) {
	defer func() {

		// recover returns the panic value
		if r := recover(); r != nil {
			fmt.Printf("worker %d recovered from panic %v\n", n, r)
			panicCounter++
			failedCounter++
		}
	}()
	for j := range ch {
		f := rand.Intn(150) + 1

		if f%2 == 0 {
			switch j.jobType {
			case "email":
				panic("Error while sending data")
			case "image":
				panic("Error while sending image")
			case "report":
				panic("Error wile sending report")
			}

		} else {
			fmt.Printf("Job %v success!\n", j.id)
			successCounter++
		}

	}

}

func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}
