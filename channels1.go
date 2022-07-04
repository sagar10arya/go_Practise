/*Understanding channels Channels Introduction
-making a channel
	-   c := make(chan int)
-putting values on a channel
	    c<-42
-taking values off of a channel
		<-c
-buffered channels : The capacity, in number of elements, set the size of buffer in the channel.
					If the capacity is 0 or absent the channel is unbuffered and communication
					succeeds only when both the sender and reciever are ready.
					c := make(chan int,1)
-unbuffered channels : If the capacity is 0 or absent the channel is unbuffered
-channels block
	-they are like runners in a relay race they are synchronized
	-they have to pass/receive the value at the same time
		-just like runners in a relay race have to pass / receive the baton to each other
		at the same time
			-one runner can't pass the baton at one moment
			-and then, later, have the other runner receive the baton
			-the baton is passed/received by the runners at the same time
		-the value is synchronously, at the same time
-channels allow us to pass values between goroutines


*/

package main

import "fmt"

func main() {

	/*c := make(chan int)
	c <- 42
		fmt.Println(<-c) */ // does not run
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)

	// c := make(chan int,1)// successful buffer

}
