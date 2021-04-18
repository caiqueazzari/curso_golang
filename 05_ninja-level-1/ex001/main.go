/*
# Hands-on Exercise #1

1. Using the short declaration operator, ASSIGN these VALUES to VARIABLES with the IDENTIFIERS "x" and "y" and "z"
  a. 42
  b. "James Bond"
  c. true

2. Now print the values stored in those variables using
  a. A single print statement
  b. Multiple print statements
*/

package main

import "fmt"

func main() {
	x := 42
	y := "James Bond"
	z := true
	fmt.Println(x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
