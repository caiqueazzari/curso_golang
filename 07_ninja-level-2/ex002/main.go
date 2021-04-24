/*
Using the following operators, write expressions and assign their values to variables:
==
<=
>=
!=
<
>
Now print each of the variables.
*/

package main

import "fmt"

func main() {
	a := (3 == 23)
	b := (3 <= 23)
	c := (3 >= 23)
	d := (3 != 23)
	e := (3 > 23)
	f := (3 < 23)
	fmt.Printf("%v\n%v\n%v\n%v\n%v\n%v\n", a, b, c, d, e, f)
}
