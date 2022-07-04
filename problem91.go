/*
In addition to the main goroutine, launch two extra go routine
	each individual goroutine should print something
use waitgroups to make sure each goroutine finishes before your program exits.
*/
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("CPUS\t\t", runtime.NumCPU())
	fmt.Println("Goroutine\t", runtime.NumGoroutine())

	wg.Add(2)
	go func() {

		fmt.Println("Hello from  queen")
		wg.Done()

	}()
	go func() {

		fmt.Println("Hello from King")
		wg.Done()

	}()

	wg.Wait()

	fmt.Println("CPUS\t\t", runtime.NumCPU())
	fmt.Println("Goroutine\t", runtime.NumGoroutine())

}
