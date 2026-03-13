package main

import (
	"fmt"
	"time"
)

// there are no classes in the golang so we use struct instead

//order struct

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // it has nanosecond precision
}

func main() {
	myorder := order{
		id:     "1",
		amount: 50,
		status: "received",
	}
	// to add the time 
	myorder.createdAt= time.Now()
	fmt.Println("order struct",myorder)

	// to get only specific field
	fmt.Println(myorder.status)
}
