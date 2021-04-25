/*
Using iota, create 4 constants for the NEXT 4 years. Print the constant values.
*/

package main

import "fmt"

const (
	a = iota + 2021
	b
	c
	d
)

func main() {
	fmt.Println(a, b, c, d)
}
