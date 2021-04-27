package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("Isso não será impresso")
	case (2 == 4):
		fmt.Println("Não funciona!!!")
	case (2 == 2):
		fmt.Println("Funciona")
		fallthrough
	case true:
		fmt.Println("Também funciona bem")
		fallthrough
	case (5 == 94):
		fmt.Println("Não funciona")
		fallthrough
	case (42 == 6):
		fmt.Println("Não funciona mesmo")
		fallthrough
	case (3 == 3):
		fmt.Println("Funciona")
	default:
		fmt.Println("Nada funcionou :(")
	}
}
