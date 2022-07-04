// WAP that adds 100 numbers to a channel
// pull the numbers of the channel and print them

package main
import "fmt"

func main(){
	c := make(chan  int)

	//send 
	go func() {
		for  i:=0 ;i<100;i++{
			c<-i
		}
		close(c)
	}()

	for v := range c{
		fmt.Println(v)
	}

	fmt.Println("about to exit")
	
}