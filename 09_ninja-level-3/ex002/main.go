/*
Print every rune code point of the uppercase alphabet three times. Your output should look like this:
65
	U+0041 'A'
	U+0041 'A'
U+0041 'A'
66
	U+0042 'B'
	U+0042 'B'
	U+0042 'B'
 … through the rest of the alphabet characters
*/

package main

import "fmt"

func main() {
	for x := 65; x < 90; x++ {
		fmt.Println(x)
		for y := 0; y < 3; y++ {
			fmt.Printf("\t%#U\n", x)
		}
	}
}
