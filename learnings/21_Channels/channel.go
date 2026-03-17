package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)
// this wil send data form main and recieve at process func (reciever goroutine)
func processNum(numChan chan int) {
	//fmt.Println("processing ...",<-numChan)
	//it prints the infinite numbers
	for num := range numChan {
		fmt.Println("Processing the number ...", num)
		time.Sleep(time.Second)
	}
}

// the communication between the go-routines is done by the channel
// it means we can send data from one go-routine to another go-routine using channel
func main() {
	//we create two separate goroutines one is process func ad other is our main func
	numChan := make(chan int)

	//sending channel
	go processNum(numChan)

	for {
		numChan <- rand.IntN(100)
	}
	//numChan <- 11

	//time.Sleep(time.Second*2)
	// messageChan := make(chan string)

	// messageChan <- "ping" // channels are blocking

	// msg := <-messageChan

	// fmt.Println(msg)
}
