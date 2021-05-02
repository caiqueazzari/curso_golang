package main

import "fmt"

func main() {
	m := map[string]int{
		"James":  32,
		"Caique": 746,
	}
	fmt.Println(m)
	delete(m, "Caique")
	fmt.Println(m)
}
