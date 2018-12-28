package main

import "fmt"

func main() {
	if true {
		fmt.Println("True")
	}

	if false {
		fmt.Println("False")
	}

	if !true {
		fmt.Println("False")
	}

	if !false {
		fmt.Println("True")
	}

	if 2 == 2 {
		fmt.Println("True")
	}

	if 2 != 2 {
		fmt.Println("False")
	}

	// Initialization, then comparitor
	if x := 42; x == 2 {
		fmt.Println("False")
	}
}
