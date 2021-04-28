/*
Create a program that uses a switch statement with no switch expression specified.
*/

package main

import "fmt"

func main() {
	x := 3
	switch {
	case x == 1:
		fmt.Println(x)
	case x == 2:
		fmt.Println(x)
	default:
		fmt.Println("default")
	}
}
