package main

import "fmt"

func main() {
	// for init; condition; post {}
	for i := 1; i <= 10; i++ {
		fmt.Printf("The outer loop: %d\n", i)
		for j := 1; j < 3; j++ {
			fmt.Printf("\t\t The inner loop: %d\n", j)
		}
	}
}
