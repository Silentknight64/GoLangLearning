package main

import "fmt"

func main() {
	x := map[string][]string{
		"bond": []string{"Miss Moneypenny", "gun", "killing bad guys"},
		"moneypenny": []string{"James Bond", "skirts", "a flower"},
	}

	for k, value := range x {
		for _, v := range value {
			fmt.Println(k, v)
		}
	}
}
