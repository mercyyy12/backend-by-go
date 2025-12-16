package main

import "fmt"

func main() {

	// defer
	defer fmt.Println("Yo bro! How i became last?")

	fmt.Println("From basic function\nThe sum is: ", add(5, 10))
	fmt.Println()

	value, err := div(10, 0)
	if err != nil {
		fmt.Println("error", err)
	} else {
		fmt.Println("The value after division is: ", value)
	}
	fmt.Println()

	s := make([]int, 0)
	s = append(s, 1, 2, 3, 4, 5) //, 6, 7, 8, 9, 10)
	count, total := sum(s)
	fmt.Printf("Named return function called\nCount = %v \nTotal = %v\n", count, total)

	// anonymous function (should be inside main or any function but not outside)
	hello := func() {
		fmt.Println("Anon function called!!")
	} // "()" here means immediate execution
	hello()

	variadic("Mango", "Apple", "Banana")

	// Closure func called
	c := closure()
	fmt.Println(c())
	fmt.Println(c())
}

// Basic function
func add(a, b int) int {
	return a + b
}

// multiple unnamed return values
func div(a, b int) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return float64(a) / float64(b), nil
}

// Named return values
func sum(a []int) (count, sum int) {
	count = 0
	sum = 0

	for i := range a {
		count++
		sum += a[i]
	}
	return
}

// Variadic function (used in APIs, logging)
func variadic(fruits ...string) {
	fmt.Println("Variadic function called")
	for i, v := range fruits {
		fmt.Println(i+1, " = ", v)
	}
}

// Closure function
func closure() func() int {
	count := 0
	return func() int {
		count++
		return count
	}

}
