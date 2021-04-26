package main

import "fmt"

func main() {
	x := 40
	if x == 40 {
		fmt.Println("x é igual a 40")
	} else if x == 41 {
		fmt.Println("x é igual a 41")
	} else {
		fmt.Println("x não é igual a 40")
	}
}
