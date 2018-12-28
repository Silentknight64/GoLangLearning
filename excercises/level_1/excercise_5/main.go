package main

import "fmt"

type person int

var x person

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)

	y := int(x)
	fmt.Println(y)
}
