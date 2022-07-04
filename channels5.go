/* Select
select statement to pull values off of multiple channels and a select statement will
pull whatever value is ready to be pulled off of that channel.
*/
package main
import "fmt"

func main(){

	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	// send
	go send(eve,odd,quit)

	// recieve

	recieve(eve,odd,quit)

	fmt.Println("About to exit")
}

func recieve(e,o,q <-chan int){
	for{
		select{
		case v := <-e:
			fmt.Println("From the eve channel  :",v)
		case v := <-o:
			fmt.Println("From the odd channel  :",v)
		case v := <-q:
			fmt.Println("From the quit channel  :",v)
			return
		}
	}
}

func send(e,o,q chan<- int){
	for i:=0 ;i<100;i++{
		if i%2 == 0{
			e<-i
		}else{
			o<-i
		}
	}
	
	q<-0
}