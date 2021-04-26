package main

import "fmt"

func main() {
	x := 0
	for {
		if x == 15 {
			fmt.Println("Feito")
			break
		}
		fmt.Println(x)
		x++
	}
}
