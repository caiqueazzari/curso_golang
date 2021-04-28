/*
Building on the previous hands-on exercise, create a program that uses “else if” and “else”.
*/

package main

import "fmt"

func main() {
	n := 20
	if n == 2021 {
		fmt.Println("Parece que n é igual a 2021 :o")

	} else if n < 2021 {
		fmt.Println("n é menor que 2021 :(")
	} else {
		fmt.Println("ELSE.")
	}
}
