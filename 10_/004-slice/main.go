package main

import "fmt"

func main() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x[2])
	fmt.Println(x)
	fmt.Println(x[1:])
	fmt.Println(x[1:3]) //slicing a slice

	for i, v := range x { //enumerate em Python
		fmt.Println(i, v)
	}
	/*
		for i := 0; i < len(x); i++ {
			fmt.Println(i, x[i])
		}
	*/
}
