package main

import "fmt"

func main() {
	i := 42
	if i == 40 {
		fmt.Println("i is 40")
	} else if i == 41 {
		fmt.Println("i is 41")
	} else if i == 42 {
		fmt.Println("i is 42")
	}
}
