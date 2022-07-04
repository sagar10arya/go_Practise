package main

import "fmt"

/* Callback
- passing a function as an argument
-functional programming not something that is recommended in go,
  however it is good to be aware of callbacks
- idiomatic go: write clear, simple, readable code

*/

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := sum(ii...) // unfurl
	fmt.Println("All numbers ", s)

	s2 := even(sum, ii...)
	fmt.Println("Even Numbers ", s2)
	s3 := odd(sum, ii...)
	fmt.Println("Odd Numbers ", s3)
}

func sum(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func even(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)

		}
	}
	return f(yi...)
}

func odd(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 != 0 {
			yi = append(yi, v)

		}
	}
	return f(yi...)
}
