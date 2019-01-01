package main

import "fmt"

type person struct {
	first 			string
	last 			string
	favoriteFlavors	[]string
}

func main() {
	p1 := person {
		first: 				"Brandon",
		last: 				"Kirk",
		favoriteFlavors:	[]string{"Chocolate", "Vanilla"},
	}

	p2 := person {
		first: 				"Miss",
		last:				"MoneyPenny",
		favoriteFlavors:	[]string{"Chocolate", "Strawberry"},
	}

	fmt.Println(p1, p2)

	for _, v := range p1.favoriteFlavors {
		fmt.Println(v)
	}

	for _, v := range p2.favoriteFlavors {
		fmt.Println(v)
	}
}

