package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	err := openFile("test.txt")
	if err != nil {

		// using error.Is
		if errors.Is(err, os.ErrNotExist) { //file does not exist
			fmt.Println("file is missing", err)
		} else if errors.Is(err, os.ErrPermission) { // permission denied
			fmt.Println("no permission", err)
		} else { // other error
			fmt.Println("other error", err)
		}

		// using error.As (in detailed form)
		var pathErr *os.PathError
		if errors.As(err, &pathErr) {
			fmt.Println("\nusing error.As:")
			fmt.Println("operation:", pathErr.Op)
			fmt.Println("file:", pathErr.Path)
			fmt.Println("root cause:", pathErr.Err)
		}

		// Using unwrapped error
		fmt.Println("\nTop level error:", err)

		// Unwrap one level
		inner := errors.Unwrap(err)
		if inner != nil {
			fmt.Println("unwrapped error:", inner)
		}

		// unwrap again
		deeper := errors.Unwrap(inner)
		if deeper != nil {
			fmt.Println("deeper error:", deeper)
		}

		// errors.Join (combine multiple errors)
		err1 := errors.New("failed to read config")
		err2 := errors.New("failed to connect to DB")

		err := errors.Join(err1, err2)
		fmt.Printf("Joined errors: %v\n", err)

		return
	}

}

func openFile(fileName string) error {
	// Opening file in write only mode
	_, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("Can't open file %w", err)
	}
	return nil
}
