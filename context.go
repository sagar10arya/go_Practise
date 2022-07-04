package main

import ("fmt"
"context")
func main(){

	// context background
	ctx := context.Background()

	fmt.Println("Context:\t",ctx)
	fmt.Println("Context err:\t",ctx.Err())
	fmt.Printf("Context type:\t%T\n",ctx)
	fmt.Println("--------------")

	// withCancel

	ctx, cancel := context.WithCancel(ctx)
	fmt.Println("Context:\t",ctx)
	fmt.Println("Context err:\t",ctx.Err())
	fmt.Printf("Context type:\t%T\n",ctx)
	fmt.Println("Cancel:\t\t",cancel)
	fmt.Printf("Cancel type:\t%T\n",cancel)
	fmt.Println("--------------")

	cancel()

	fmt.Println("Context:\t",ctx)
	fmt.Println("Context err:\t",ctx.Err())
	fmt.Printf("Context type:\t%T\n",ctx)
	fmt.Println("Cancel:\t\t",cancel)
	fmt.Printf("Cancel type:\t%T\n",cancel)
	fmt.Println("--------------")

}