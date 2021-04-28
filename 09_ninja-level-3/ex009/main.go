/*
Create a program that uses a switch statement with the switch expression specified as a variable of TYPE string with the IDENTIFIER “favSport”.
*/

package main

import "fmt"

func main() {
	favSport := "Swimming"
	switch favSport {
	case "Tenis":
		fmt.Println("My favorite sport is Tenix")
	case "Swimming":
                fmt.Println("My favorite sport is Swimming")
	case "Soccer":
                fmt.Println("My favorite sport is Soccer")
	}
}
