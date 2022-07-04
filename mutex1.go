package main

import (
	"fmt"
	"runtime"
	"sync"
)

// 1 race condition
// we will use mutex to avoid race condition

func main() {
	fmt.Println("CPUs:\t", runtime.NumCPU())
	fmt.Println("Goroutine:\t", runtime.NumGoroutine())

	counter := 0
	const gs = 100

	var wg sync.WaitGroup
	wg.Add(gs)

	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()

			wg.Done()

		}()
		fmt.Println("Goroutine:\t", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutine:\t", runtime.NumGoroutine())
	fmt.Println("count: ", counter)

}
