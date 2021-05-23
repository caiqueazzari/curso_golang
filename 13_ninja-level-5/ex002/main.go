/*
Take the code from the previous exercise, then store the values of type person in a map with the key of last name. Access each value in the map. Print out the values, ranging over the slice.
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

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, v := range m {
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, val := range v.favoriteFlavors {
			fmt.Println(i, val)
		}
		fmt.Println("-------")
	}
}
