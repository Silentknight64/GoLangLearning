package main

import "fmt"

// Typed constants
const (
	a int = 3
	b bool = true
	c string = "James Bond"
)

// Untyped constants
const (
	d = 5
	e = false
	f = "Pizza"
)

func main() {
	fmt.Println(a, b, c, d, e, f)
}
