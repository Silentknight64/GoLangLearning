package _2_var

import "fmt"

// Global variable
var y = 43

// Assigns zero value of type int to "z"
var z int

func main() {
	x := 42
	fmt.Println(x)

	fmt.Println(y)

	foo()
}

func foo() {
	fmt.Println(y)
}
