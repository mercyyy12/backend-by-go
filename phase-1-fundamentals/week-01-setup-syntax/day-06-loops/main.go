package main

import "fmt"

func main() {
	// for loop with init and condition
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	x := 0

	// while style for loop
	// for x < 10 {
	// 	fmt.Println(x)
	// 	x++
	// }

	// infinite loop unless break
	for {
		if x == 5 {

			break
		}
		fmt.Println(x)
		x++
	}

	// nested loop with labeled break
	// outer:
	// 	for i := 1; i <= 3; i++ {
	// 		for j := 1; j <= 3; j++ {
	// 			if i == 2 && j == 2 {
	// 				fmt.Println("Breaking outer loop")
	// 				break outer
	// 			}
	// 			fmt.Printf("i=%d j=%d\n", i, j)
	// 		}
	// 	}

	// for-range loop over string
	v := "Mercyyy"
	for i, v := range v {
		fmt.Println(i, string(v))
	}

	// for i, r := range "नमस्ते" {
	// 	fmt.Println(i, string(r))
	// }

}
