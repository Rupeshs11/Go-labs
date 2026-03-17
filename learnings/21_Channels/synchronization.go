package main

import "fmt"


// go routine synchronizer
func task(done chan bool) {
	defer func(){ done <- true}()
	fmt.Println("processing...")
	
}

func main(){

	done := make(chan bool)
	go task(done)
	<-done // block
}