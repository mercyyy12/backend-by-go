package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Reads string input
func input() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strings.ToLower(strings.TrimSpace(text))
}

// Reads integer input
func inputInt() int {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	num, err := strconv.Atoi(strings.TrimSpace(text))
	if err != nil {
		fmt.Println("Please enter a integer")
		return -1
	}
	return num
}
