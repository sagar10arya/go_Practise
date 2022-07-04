/* Create a slice of a slice of a string ([][]string) Store the data in
multidimensional slice :
"James","Bond","Shaken, not stirred"
"Miss","Moneypenny","Hello james"*/

package main

import "fmt"

func main() {
	xs1 := []string{"James", "Bond", "Shaken, not stirred"}
	xs2 := []string{"Miss", "Moneypenny", "Hello james"}
	fmt.Println(xs1)
	fmt.Println(xs2)

	xxs := [][]string{xs1, xs1}
	fmt.Println(xxs)

	for i, xs := range xxs {
		fmt.Println("record: ", i)
		for j, val := range xs {
			fmt.Printf("\t index position: %v \t value : %v\n", j, val)
		}
	}
}
