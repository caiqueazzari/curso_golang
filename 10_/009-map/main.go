package main

import "fmt"

func main() {
	m := map[string]int{
		"James":  32,
		"Caique": 64,
		"Lucas":  58,
	}
	fmt.Println(m)
	fmt.Println(m["Caique"])    // will return 64
	fmt.Println(m["sadasiods"]) //will return the zero value
	v, ok := m["sadasiods"]
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["sadasiods"]; ok {
		fmt.Println("This is the if print", v)
	}
}
