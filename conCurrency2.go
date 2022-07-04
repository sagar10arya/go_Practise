package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs:\t", runtime.NumCPU())
	fmt.Println("Goroutine:\t", runtime.NumGoroutine())

	counter := 0
	const gs = 100

	var wg sync.WaitGroup
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()

		}()
		fmt.Println("Goroutine:\t", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutine:\t", runtime.NumGoroutine())
	fmt.Println("count: ", counter)

	// 1 race condition
	// will use mutex to avoid race condition
}
