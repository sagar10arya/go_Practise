// Write a program that prints a number in decimal,binary,hex

package main

import "fmt"

func main() {
	x := 42

	fmt.Printf("Number in decimal: %d\n", x)
	fmt.Printf("Number in binary: %b\n", x)
	fmt.Printf("Number in hexa: %#x\n", x)
}
