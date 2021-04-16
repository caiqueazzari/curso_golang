package main

import "fmt"

func main() {
	fmt.Println("Hello! My cat Piki is very cute")
	potato()

	for i:=0; i < 10; i++ {
		if i % 2 == 0{
			fmt.Println(i)
		}
	}
}

func potato() {
	fmt.Println("\nI like potatoes\n ")
}