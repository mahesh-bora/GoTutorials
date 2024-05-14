package main

import "fmt"


// Channels in Go are a powerful feature for communication and synchronization between goroutines 
func main(){
	var c = make(chan int, 10) //Buffer of 10 so that the process function can add upto 10 values to the buffer
	go process(c)
	for i:= range c{
		fmt.Println(i)
	}
}

func process(c chan int){
	defer close(c) // Need this to notify the main function that the process is ready to terminate
	// Deadlock happens if the process is not closed
	for i := 0; i<5; i++{
		c<-i
	}
	fmt.Println("Exiting the process")
}

