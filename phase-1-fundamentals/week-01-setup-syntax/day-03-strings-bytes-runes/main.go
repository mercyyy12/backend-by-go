package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Type conversion
	x := 65
	// var t float32 = x //implicit conversion is not allowed
	// fmt.Println(t)

	// Explicit conversion
	// int to float (float to int = vice versa)
	var f float64 = float64(x)
	// or (both are same)
	// f := float64(x)
	fmt.Printf("The value of f is %v and it's type is %T\n", f, f)

	// int to string
	var alpha string = strconv.Itoa(x)
	fmt.Printf("The value of s is %v and it's type is %T\n", alpha, alpha)

	// string to int
	s_age := "21"
	age, err := strconv.Atoi(s_age)
	if err != nil {
		fmt.Println("invalid number")
	}
	fmt.Printf("The value of age is %v and it's type is %T\n", age, age)

	// string indexing and immutability
	var s string = "hello"
	fmt.Printf("%v\n", s[0]) // 104
	fmt.Printf("%c\n", s[0]) // h

	// s[0] = 'a' ERROR strings are immutable
	// fmt.Printf("%c", s[0])

	// byte = uint8
	var b byte = 'M'
	//var c byte = 'Me' ERROR should contain only one byte
	fmt.Printf("byte: %v (%T)\n", b, b)

	// rune = int32 (Unicode code point)
	var r rune = '🔥'
	fmt.Printf("rune: %v (%T)\n", r, r)
}
