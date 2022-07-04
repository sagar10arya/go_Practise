/* Create a slice to store the names of all the states in USA. Use make and
append to do this.
Goal: Don not have the array that underlies the slice created more than once.
What is the length of the slice? What is the capacity of the slice? Print out
all the values with their index positions, without using the range clause.
Here is  the listof all 50 states:
` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`,
` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`,
 ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`,
  ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`,
   ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`,
   ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`,
    ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`,
	` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`,
	 ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`,
	  ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`} */

package main

import "fmt"

func main() {

	// slice_name := make([]type, length, capacity)

	slice1 := make([]string, 50, 50)
	fmt.Println("First Time:")
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	states := []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`,
		` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`,
		` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`,
		` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`,
		` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`,
		` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`,
		` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`,
		` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`,
		` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`,
		` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`,
		` West Virginia`, ` Wisconsin`, ` Wyoming`}

	slice1 = append(slice1, states...)

	fmt.Println("Second Time:")
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	for i := 0; i < len(states); i++ {
		fmt.Println(i, states[i])
	}

}
