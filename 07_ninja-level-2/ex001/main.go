/*
Write a program that prints a number in decimal, binary, and hex
*/

package main

import "fmt"

func main() {
	x := 10
	fmt.Printf("%d\n%b\n%#x\n", x, x, x)
}
