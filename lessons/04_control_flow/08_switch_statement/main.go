package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("This should not print")
	case 2 == 4:
		fmt.Println("this should not print")
	case 3 == 3:
		fmt.Println("prints")
		fallthrough
	case 4 == 4:
		fmt.Println("Also true, does it print?")
	default:
		fmt.Println("This is default")
	}

	switch n := "Bond"; n {
	case "James":
		fmt.Println("He's james!")
	case "Bond":
		fmt.Println("He's Bond!")
	default:
		fmt.Println("He's neither James nor Bond.")
	}

}
