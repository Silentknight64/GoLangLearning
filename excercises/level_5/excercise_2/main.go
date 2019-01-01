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

	people := map[string]person {
		p1.last: p1,
		p2.last: p2,
	}

	for key, value := range people {
		for _, v := range value.favoriteFlavors {
			fmt.Println(key, v)
		}
	}
}

