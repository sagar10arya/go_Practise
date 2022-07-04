package main

import "fmt"

/* Closure
- one scope enclosing other scope
	variable declared in the outer scope are accessiblein inner scopes
- closure helps us limit the scope of variables.

{

} -- used for limit of scope
*/
func main() {
	a := incrementor()
	b := incrementor()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())

}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
