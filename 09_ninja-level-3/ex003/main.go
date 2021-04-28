/*
Create a for loop using this syntax
for condition { }
Have it print out the years you have been alive.
*/

package main

import "fmt"

func main() {
	nasc := 1930
	for nasc <= 2021 {
		fmt.Println(nasc)
		nasc += 1
	}
}
