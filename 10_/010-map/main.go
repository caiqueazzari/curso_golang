package main

import "fmt"

func main() {
	m := map[string]int{
		"James":  32,
		"Caique": 746,
	}

	m["Todd"] = 33

	for k, v := range m {
		fmt.Println(k, v)
	}

	xi := []int{4, 5, 6, 7, 4, 6}

	for i, v := range xi {
		fmt.Println(i, v)
	}
}
