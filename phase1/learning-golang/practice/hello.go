package main

import (
	"bufio"
	"fmt"
	"go/scanner"
	"os"
	"strconv"
	"strings"
)

func main() {
	// fmt.Printf("Enter your age: ")
	// reader := bufio.NewReader(os.Stdin)
	// input, err := reader.ReadString('\n')
	// if err != nil {
	// 	fmt.Println("error occured", err)
	// 	return
	// }
	// age, err := strconv.Atoi(strings.TrimSpace(input))
	// if err != nil {
	// 	fmt.Println("error occured", err)
	// 	return
	// }

	scanner := bufio.NewReader(os.Stdin)
	input := scanner.ReadString("")

}
