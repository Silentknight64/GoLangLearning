package main

import "fmt"

func main() {
	m := map[string]int{
		"James": 			42,
		"Miss Moneypenny": 	27,
	}
	fmt.Println(m)
	fmt.Println(m["James"])
	fmt.Println(m["Barnabas"])

	v, ok := m["Barnabas"]
	fmt.Println(v, ok)

	// Adding a value to the map
	m["todd"] = 33

	if v, ok := m["Miss Moneypenny"]; ok {
		fmt.Println(v)
	}

	for k, v := range m {
		fmt.Println(k, v)
	}

	// Delete a value from the map
	delete(m, "James")
	fmt.Println(m)

	if v, ok := m["Miss Moneypenny"]; ok {
		fmt.Println("Value: ", v)
		delete(m, "Miss Moneypenny")
	}
}
