// Range 
// during the range of channel we also tell about to close the channel

package main
import "fmt"
func main(){
	c := make(chan int)

	//send
go func (){
	for i:=0;i<5;i++{
		c <- i
	}
	close(c)
}()
	// recieve

	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("About to exit")
}



