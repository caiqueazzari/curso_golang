/*Using the code from the previous example, delete a record from your map. Now print the map out using the “range” loop
 */

package main

import "fmt"

func main() {
	m := map[string][]string{
		"a_cc": []string{"a", "b", "c"},
		"f_i":  []string{"a", "b", "b"},
		"a_a":  []string{"a", "b", "c"},
	}

	m["x_y"] = []string{"a", "b", "c"}

	delete(m, "a_a")

	for k, v := range m {
		fmt.Println(k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}
