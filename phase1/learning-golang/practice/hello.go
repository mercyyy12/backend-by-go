// range loops
//
// # Error handling
//
// defer statement
//
// Functions with error returns
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter 10 elements: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error: \n", err)
	}
	input = strings.TrimSpace(input)
	input, err = strconv.Atoi(input)
	parts := strings.Split(input, " ")

	// if err != nil {
	// 	fmt.Printf("Error: \n", err)
	// }
	sum := 0
	for i := 0; i < len(input)-1; i++ {

		sum += parts[i] + parts[i+1]
	}
	fmt.Printf("the sum is: %d", sum)
}
