/*Create a map of key of type string whic is person's last_first name and value of
type strings store their favorite things. Store 3 records on your map. Print out
all the values with their index positions in the slice, without usig the range clause.
*/

package main

import "fmt"

func main() {

	m := map[string][]string{
		`bond james`:      []string{`shaken, not stirred`, `martinis`, `women`},
		`moneypenny_miss`: []string{`james bond`, `literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `ice-cream`, `sunset`},
	}
	// fmt.Println(m)
	for k, v := range m {
		fmt.Println("This is the record for ", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}

}
