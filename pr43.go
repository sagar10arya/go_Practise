// Using the code from pr2, Use slicing to create the following new slices
// which are printed as:
// [42,43,44,45,46]
// [47,48,49,50,51]
// [44,45,46,47,48]
// [43,44,45,46,47]

package main

import "fmt"

func main() {

	s1 := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Println(s1[:5])
	fmt.Println(s1[5:])
	fmt.Println(s1[2:7])
	fmt.Println(s1[1:6])

}
