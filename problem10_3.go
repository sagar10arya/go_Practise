// using code, pull the values off the channel using a range loop

package main
import "fmt"

func main(){
	c := gen()
	recieve(c)

	fmt.Println("About to exit")
}

func recieve(c <-chan int){
	for v:= range c{
		fmt.Println(v)
	}
}

func gen() <-chan int{
	c := make(chan int)

	go func(){
		for i:=0 ;i<10 ;i++{
			c <- i
		}
		close(c)
	}()
	return c
}