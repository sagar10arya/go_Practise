// USe the defer keyword to shoe that a deferred func runs after the func containig itself

package main

import "fmt"

func main() {
	defer foo()
	bar()
}

func foo() {
	fmt.Println("I am deferred and will run at last ")
}
func bar() {
	fmt.Println("Hey I am bar")
}
