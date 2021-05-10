package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "Caique",
		last:  "Azzari",
		age:   42,
	}

	p2 := person{
		first: "Iris",
		last:  "F",
		age:   67,
	}
	fmt.Println(p1)
	fmt.Println(p1)

	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)
}
