package main

import "fmt"

type vehicle struct {
	doors		int
	color 		string
}

type truck struct {
	vehicle
	fourWheel 	bool
}

type sedan struct {
	vehicle
	luxury 		bool
}

func main() {
	truck := truck {
		vehicle: vehicle {
			doors: 2,
			color: "Blue",
		},
		fourWheel: true,
	}

	sedan := sedan {
		vehicle: vehicle {
			doors: 4,
			color: "Red",
		},
		luxury: true,
	}

	fmt.Println(truck)
	fmt.Println(sedan)

	fmt.Println(truck.doors, truck.color, truck.fourWheel)
	fmt.Println(sedan.doors, sedan.color, sedan.luxury)
}