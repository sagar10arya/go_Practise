/* Directional channels :
 you could send on to it and you can recieve from it. But you couls also specify 
 a channel is only one direction which means you can only send to this channel
 and toy could only recieve from this channel.
*/
package main
import "fmt"
func main() {
/*	// c := make(chan int, 2)  // bidirectional	
	c := make(chan <- int, 2)  // only a send channel
	c := make(<-chan int, 2)  // only a recieve channel

	c <- 42
	c <- 43

	fmt.Println(<-c) 
 // invalid operation: cannot receive from send-only channel c (variable of type chan<- int)
	fmt.Println(<-c)
	*/ 		// Doesn't work

	c := make(chan int)  // bidirectional	
	cs := make(chan <- int)  // only a send channel
	cr := make(<-chan int)  // only a recieve channel


	fmt.Println("-----------")
	fmt.Printf("c\t%T\n",c)
	fmt.Printf("cr\t%T\n",cr)
	fmt.Printf("cs\t%T\n",cs)

	// general to specific assigns
	// cr = c
	// cs = c

	fmt.Printf("c\t%T\n",(<-chan int)(c))
	fmt.Printf("c\t%T\n",(chan <- int)(c))

/* A channel may  be  constrained only to send or only to recieve[general to specific] 
by conversion or assignment
dosen't assigns - specific to general and specific to specific
assigns - general to specific
conversion : general to specific works */

}