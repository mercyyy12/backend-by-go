package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// bufio reader to take string input
func input() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return "thik xaina", fmt.Errorf("Input lida erro %w", err)
	}
	text = strings.TrimSpace(text)
	return text, nil
}

// bufio reader to take integer input
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

// closure func to generate IDs
func countId() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}
