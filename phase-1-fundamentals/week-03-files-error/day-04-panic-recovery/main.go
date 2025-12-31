package main

import "fmt"

func main() {

	// fmt.Println("")

	a, b := 5, 5
	result, err := div(a, b)
	if err != nil {
		fmt.Println("error occured")
	}
	fmt.Println("a divided by b =", result)

	for i := 0; i < 5; i++ {
		func(n int) {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered from panic in iteration", n, ":", r)
				}
			}()
			process(n)
		}(i)
	}

}

func div(a, b int) (int, error) {
	if b == 0 {
		// panic stops the program immediately
		panic("Can't divide with 0")
	}
	return a / b, nil
}

func process(i int) {
	if i == 3 {
		panic("panic occureddddd")
	}
	fmt.Println("process completed = ", i)

}
