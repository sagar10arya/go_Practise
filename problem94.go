// Fix the race condition you created in the problem93.go using mutex
// it makes sense to remove runtime.Gosched
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs:\t\t", runtime.NumCPU())
	fmt.Println("Goroutine:\t", runtime.NumGoroutine())

	incrementer := 0
	const gs = 100

	var wg sync.WaitGroup
	wg.Add(gs)

	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := incrementer
			v++
			incrementer = v
			fmt.Println("Incrementer", incrementer)
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end value: ", incrementer)

	fmt.Println("CPUs:\t\t", runtime.NumCPU())
	fmt.Println("Goroutine:\t", runtime.NumGoroutine())
}
