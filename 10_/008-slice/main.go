package main

import "fmt"

func main() {
	ca := []string{"Caique", "Azzari", "Chocolate", "Morango"}
	fmt.Println(ca)
	iri := []string{"Iris", "Flores", "Flocos", "Creme"}
	fmt.Println(iri)
	//multi-dimensional slice
	xp := [][]string{ca, iri} //"slice of slice of string" slice of people
	fmt.Println(xp)           // slice of slice of string
}
