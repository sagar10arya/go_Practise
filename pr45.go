/* To delete from a slice wew use Append along with the slicing.
Starts with the slice 	x := []int{42,43,44,45,46,47,48,49,50,51}
use append and slicing to get the values here which you should print then
[42,43,44,48,49,50,51] */

package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x[:2], x[6:]...)

	fmt.Println(x)

}
