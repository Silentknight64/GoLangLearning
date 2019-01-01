package main

import "fmt"

func main() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(len(x))
	fmt.Println(x)
	fmt.Println(x[0])

	for i, v := range x {
		fmt.Println(i, v)
	}

	fmt.Println(x[1:])
	fmt.Println(x[1:3])

	x = append(x, 77, 88, 99, 1014)
	fmt.Println(x)

	y := []int{235, 456, 678, 987}
	x = append(x, y...)

	// Deleting items at index 2 and 3
	x = append(x[:2], x[4:]...)

	for _, v := range x {
		fmt.Println(v)
	}

	// len 10 cap 100
	x = make([]int, 10, 100)

	jb := []string{"James", "Bond", "Chocolate", "Martini"}
	fmt.Println(jb)
	mp := []string{"Miss", "Moneypenny", "Strawberry", "Hazelnut"}
	fmt.Println(mp)

	// Multi-dimensional slice
	xp := [][]string{jb, mp}
	fmt.Println(xp)
}
