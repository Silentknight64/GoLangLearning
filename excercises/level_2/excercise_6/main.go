package main

import "fmt"

// Iota for years
const (
	a = iota + 2014
	b
	c
	d
)

func main() {
	fmt.Println(a, b, c, d)
}
