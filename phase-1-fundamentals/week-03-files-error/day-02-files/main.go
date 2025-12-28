package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	// Opening file
	// file, err := os.Open("test.txt")
	// if err != nil {
	// 	fmt.Println("Error opening file", err)
	// 	return
	// }
	// fmt.Println(file)
	// defer file.Close()

	// Creating a file
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Error opening file", err)
		return
	} else {
		fmt.Println("File created successfully!")
	}
	defer file.Close()

	// writing in a file
	file.WriteString("Hello everyone,\n")
	n, err := file.WriteString("My name is Mercyyy\n")

	if err != nil {
		fmt.Println("Write error:", err)
		return
	}
	fmt.Println("Bytes written:", n)

	// Reading from file (entirely)
	data, err := os.ReadFile("test.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("File content:\n", string(data))

	// Reading line by line using bufio reader
	data1, err := os.Open("test.txt")
	reader := bufio.NewReader(data1)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Print(line) // last line (may not end with \n)
				break
			}
			fmt.Println("Read error:", err)
			return
		}
		fmt.Print(line)
	}

	// Appending to a file

	// 0644 = file permission (r/w for owner, read only for others)
	// 0755 = file permission (r/w/ex for owner, r/ex for others)
	// os.O_APPEND = Add content at the end
	// os.O_WRONLY = Write only mode

	file2, err := os.OpenFile("test.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file2.Close()

	_, err = file2.WriteString("I am appended!\n")
	if err != nil {
		fmt.Println("Write error:", err)
	}

	// common file errors
	file3, err := os.Open("test.txt")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File does not exist")
		} else if os.IsPermission(err) {
			fmt.Println("Permission denied")
		} else {
			fmt.Println("Other error:", err)
		}
		return
	}
	defer file3.Close()

	// // Removing file
	// err = os.Remove("test.txt")
	// if err != nil {
	// 	fmt.Println("Error deleting file:", err)
	// } else {
	// 	fmt.Println("File deleted successfully")
	// }

}
