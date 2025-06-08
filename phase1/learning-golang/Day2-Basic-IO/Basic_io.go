package main

func main() {
	// fmt basic input examples

	/*
		var age int
		var health, name string

		fmt.Print("Enter your age: ") //prints the text to the console without a newline
		fmt.Scan(&age) // Reads input until space or newline

		fmt.Println("How are you feeling today? ") //prints the text to the console with a newline at the end
		fmt.Scanln(&health) // reads input until newline

		fmt.Print("What is your name? ")
		fmt.Scanln(&name)

		fmt.Printf("Hello %s, your age is %d and you're feeling %s today.\n", name, age, health) //prints formatted output

		// Example of fmt.Scanf
		var day int
		var month string
		fmt.Print("Enter your birthday day and month: ")
		fmt.Scanf("%d %s", &day, &month) //reads formatted input according to the format specifier
		fmt.Printf("Your birthday is on %d of %s.\n", day, month)
	*/

	// bufio.Reader (stdin)

	/*
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter your age: ")
		input, err := reader.ReadString('\n') // reads until newline
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		input = strings.TrimSpace(input) // clean up space
		age, err := strconv.Atoi(input)  // convert string to int
		if err != nil {
			fmt.Println("Invalid input! Please enter a number.")
			return
		}
		fmt.Println("You are", age, "years old.")
	*/

	//bufio.Reader (file input)

	/*
		file, err := os.Open("data.txt") // open existing file
		if err != nil {
			fmt.Println("error in file", err)
			return
		}
		defer file.Close()

		reader := bufio.NewReader(file)
		for {
			line, err := reader.ReadString('\n') // read line by line
			if err != nil {
				break
			}
			fmt.Print("Line: ", line)
		}
	*/

	// bufio.Scanner (stdin)

	/*
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Enter your age: ")
		if !scanner.Scan() {
			fmt.Println("Failed to read input.")
			return
		}
		input := scanner.Text()
		input = strings.TrimSpace(input)
		age, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid age!")
			return
		}
		fmt.Println("You are", age, "years old.")
	*/

	// bufio.Scanner (file input)

	/*
		file, err := os.Open("data.txt") // open file for reading
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			fmt.Println("Line:", scanner.Text()) // print each line
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Scan error:", err)
		}
	*/

	//bufio.Writer (write to file)

	/*
	   file, err := os.Create("output.txt") // create a new file

	   	if err != nil {
	   		fmt.Println("Error creating file:", err)
	   		return
	   	}

	   defer file.Close() // close file when function ends

	   writer := bufio.NewWriter(file)

	   // Write multiple lines to the buffer
	   writer.WriteString("Line 1: Hello, World!\n")
	   writer.WriteString("Line 2: This is Go.\n")
	   writer.WriteString("Line 3: Writing to a file.\n")
	   writer.WriteString("Line 4: bufio.Writer is fast!\n")

	   writer.Flush() // flush buffer to file
	*/
}
