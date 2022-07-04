// WAP that launches 10 goroutines. Each goroutine adds 10 numbers to a channel.
// pull 100 numbers to a channel and print them.

package main
import "fmt"

func main(){

	c := make(chan  int)

	for j:=0; j<10; j++{
		go func() {
			for  i:=0 ;i<100;i++{
				c<-i
			}
			close(c)
			
		}()
	}
	

	for k:=0 ;k<100 ;k++{
		fmt.Println(k,<-c)
	}

	fmt.Println("about to exit")
}