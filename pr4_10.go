// Using the code from probelm 9 delete a record from the map.

package main

import "fmt"

func main() {

	m := map[string][]string{
		`bond james`:      []string{`shaken, not stirred`, `martinis`, `women`},
		`moneypenny_miss`: []string{`james bond`, `literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `ice-cream`, `sunset`},
	}
	// fmt.Println(m)

	m[`fleming_ian`] = []string{`steaks`, `cigars`, `espionage`}
	delete(m, `no_dr`)
	for k, v := range m {
		fmt.Println("This is the record for ", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}

}
