package main

import "fmt"

func main() {
	x := map[string][]string{
		"bond": []string{"Miss Moneypenny", "gun", "killing bad guys"},
		"moneypenny": []string{"James Bond", "skirts", "a flower"},
	}

	x["kirk"] = []string{"programming", "girls", "puppies"}

	delete(x, "bond")

	for k, v := range x {
		for _, value := range v {
			fmt.Println(k, value)
		}
	}
}
