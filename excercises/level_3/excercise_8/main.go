package main

import "fmt"

func main() {
	switch {
	case 2 != 2:
		fmt.Println("False")
	case 4 == 4:
		fmt.Println("True")
	default:
		fmt.Println("Default")
	}
}
