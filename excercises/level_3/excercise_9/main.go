package main

import "fmt"

func main() {
	switch favSport := "Soccer"; favSport {
	case "Football":
		fmt.Println("Football")
	case "Soccer":
		fmt.Println("Soccer")
	default:
		fmt.Println("Default")
	}
}
