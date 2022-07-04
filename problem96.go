// Create a program that prints your OS and Arch.Use the following commands to run it:
// go run ,  go build, go install

package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("OS: ", runtime.GOOS)
	fmt.Println("Architecture: ", runtime.GOARCH)
}
