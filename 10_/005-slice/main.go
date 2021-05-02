package main

import "fmt"

func main() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)
	x = append(x, 77, 88, 99, 34, 6564, 45, 3, 4)
	fmt.Println(x)

	y := []int{4, 6, 3, 77, 5, 243}
	x = append(x, y...)
	fmt.Println(x)
}
