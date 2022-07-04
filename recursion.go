package main

import "fmt"

func main() {
	a := factorial(4)
	fmt.Println("factorial using recursion ", a)

	fmt.Println("factorial using loops ", loopfact())
}

// recursion have negative effect of memory usage.
// This can be done using lopos which is best option I guess.
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func loopfact() int {
	// Using loops
	fact := 1
	n := 4
	for i := 1; i <= n; i++ {
		fact = fact * i
	}
	return fact
}
