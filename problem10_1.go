// get the code working using func literal, aka, anonymous self executing func
// a buffered channel

package main
import "fmt"

func main(){
	// c := make(chan int)
	c := make(chan int,1) // buffered

	/* go func() {
		c <- 42
	}() */ // anonymous self executing func 

	c<-42
	fmt.Println(<-c)
}