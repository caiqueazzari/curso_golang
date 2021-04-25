/*
Create TYPED and UNTYPED constants. Print the values of the constants.
*/

package main

import "fmt"

const (
	a         = 12
	b         = 32
	c int     = 45
	d float64 = 54.35
	e string  = "Caique Azzari"
)

func main() {
	fmt.Println("Untyped constants:\n")
	fmt.Printf("a: %v \ttype: %T\n", a, a)
	fmt.Printf("b: %v \ttype: %T\n", b, b)
	fmt.Println("\nTyped constants:\n")
	fmt.Printf("c: %v \ttype: %T\n", c, c)
	fmt.Printf("d: %v \ttype: %T\n", d, d)
	fmt.Printf("e: %v \ttype: %T\n", e, e)
}
