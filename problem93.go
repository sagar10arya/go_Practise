/*
-Using goroutines, create an incrementer program value
	-have a variable to hold the incrementer
    -launch a bunch of goroutines
		-each goroutine should
			-read the incrementer value
				-store it in a new variable
			-yield the processor with runtime. Gosched() increment the new variable
			-write the value in the new variable back to the incrementer variable
-use waitgroups to wait for all of your goroutines to finish
-the above will create a race condition
-Prove that it is a race condition by using the -race flag

*/
// Gosched yields the processor, allowing other goroutines to work.
// It does not suspend the other  goroutine, so execution resumes automatically

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs:\t", runtime.NumCPU())
	fmt.Println("Goroutine:\t", runtime.NumGoroutine())

	incrementer := 0
	const gs = 100

	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := incrementer
			runtime.Gosched()
			v++
			incrementer = v
			fmt.Println("Incrementer", incrementer)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end value: ", incrementer)
}
