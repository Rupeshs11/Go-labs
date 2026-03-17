package main

import (
	"fmt"
	"time"
)

func emaiSender(emailChan chan string,done chan bool){
	defer func(){done <- true}()
	for email := range emailChan{
		fmt.Println("sending email to",email)
		time.Sleep(time.Second)
	}

}

func main() {
	emailChan := make(chan string, 100)
	//synchonizer channel
	done := make(chan bool)

	go emaiSender(emailChan,done)

	for i:=0;i<5;i++{
		emailChan <- fmt.Sprintf("%d@gmail.com",i)
	}

	fmt.Println("done sending.")

	// emailChan <- "knox@gmail.com"
	// emailChan <- "luffy@gmail.com"

	// fmt.Println(<-emailChan)
	// fmt.Println(<-emailChan)
	// we need to close(emailChan) to avoid the deadlock after last receiving the last request
	close(emailChan)
	<- done
}
