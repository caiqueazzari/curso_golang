/*
Create your own type “person” which will have an underlying type of “struct” so that it can store the following data:
first name
last name
favorite ice cream flavors
Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice which stores the favorite flavors.
*/

package main

import "fmt"

type person struct {
	first           string
	last            string
	favoriteFlavors []string
}

func main() {
	p1 := person{
		first: "Caique",
		last:  "Azzari",
		favoriteFlavors: []string{
			"chocolate",
			"baunilha",
		},
	}

	p2 := person{
		first: "Iris",
		last:  "F",
		favoriteFlavors: []string{
			"flocos",
			"baunilha",
			"limão",
		},
	}

	fmt.Println(p1.first)
	fmt.Println(p1.last)
	for i, v := range p1.favoriteFlavors {
		fmt.Println(i, v)
	}

	fmt.Println(p2.first)
	fmt.Println(p2.last)
	for i, v := range p2.favoriteFlavors {
		fmt.Println(i, v)
	}
}
