/*
Waitgroup --> A waitgroup allows us to wait until a certain process is finished.
Important Concepts :-
	-sync.Waitgroup
	-runtime.NumCPU()
	-runtime.NumGoroutine()

If there is 1 CPU in computer and launch the goroutine than that function go into goroutine,
 but function still continuous to run and other function will going on after completing own
 goroutine than come back to other goroutine before the completing the this goroutine all the
 main func will completed and program exit.



In wait group:

Step 1: Define the var like
                   var wg sync.WaitGroup
Step 2: wg.Add(1)                  //How group is waiting
Step 3: wg.Wait()                    // It's say wait here some function waiting and check it.
Step 4: wg.Done()                  // Write that function after completing the function.
To synchronise(Put into wait) any function  just use "go" front of the function.
Or launching into go routine.
go foo()
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	fmt.Println("OS\t\t", runtime.GOOS)                // name of the OS
	fmt.Println("ARCH\t\t", runtime.GOARCH)            // Architecture
	fmt.Println("CPUS\t\t", runtime.NumCPU())          // No. of CPUs
	fmt.Println("Goroutine\t", runtime.NumGoroutine()) // No. of Goroutines

	wg.Add(1)
	go foo()
	bar()

	fmt.Println("CPUS\t\t", runtime.NumCPU())
	fmt.Println("Goroutine\t", runtime.NumGoroutine())
	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done()
}
func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}
