package main

import "fmt"

func main() {

	a, b := 5, 10
	result, err := div(a, b)
	if err != nil {
		fmt.Println("error occured")
	}
	fmt.Println("a divided by b =", result)

	for i := 0; i < 5; i++ {

		// Outer anon func
		func(n int) {

			// Inner anon func with defer to catch panic
			defer func() {

				// recover returns the panic value
				if r := recover(); r != nil {
					fmt.Println("Recovered from panic in iteration", n, ":", r)
				}
			}()

			// this runs first and then the defer func
			process(n)

		}(i) // Immediate execution if outer anon func
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
