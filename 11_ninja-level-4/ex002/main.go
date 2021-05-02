/*
Using a COMPOSITE LITERAL: 
create a SLICE of TYPE int
assign 10 VALUES 
Range over the slice and print the values out. 
Using format printing
print out the TYPE of the slice
*/

package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 44, 53, 65, 77, 87, 99, 100}
	for _, v := range x {
		fmt.Println(v)
	}
	fmt.Printf("%T\n", x)
}
