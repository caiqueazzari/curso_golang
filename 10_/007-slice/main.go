package main

import "fmt"

func main() {
	x := make([]int, 10, 100) // make é útil quando sabemos o tamanho máximo do array.
	fmt.Println(x)
	fmt.Println(len(x)) //tamanho
	fmt.Println(cap(x)) //capacidade
	x[0] = 47
	x[9] = 2392
	fmt.Println(x)
	x = append(x, 432432423)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}
