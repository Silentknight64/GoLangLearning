package main

import "fmt"

func main() {
	x := 1

	// Condition only
	for x < 10 {
		fmt.Println(x)
		x++
	}

	// Infinite loop
	for {
		if x < 9 {
			break
		}
		fmt.Println(x)
		x++
	}
}
