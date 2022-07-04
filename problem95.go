// // Fix the race condition you created in the problem94.go using atomic

package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPUs:\t", runtime.NumCPU())
	fmt.Println("Goroutine:\t", runtime.NumGoroutine())

	var incrementer int64
	const gs = 100

	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&incrementer, 1)
			runtime.Gosched()

			fmt.Println("Incrementer\t", atomic.LoadInt64(&incrementer))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end value: ", incrementer)
	fmt.Println("Goroutine:\t\t", runtime.NumGoroutine())

}
