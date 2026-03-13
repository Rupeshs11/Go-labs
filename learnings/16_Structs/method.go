package main

import (
	"fmt"
	"time"
)

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
}

// we will write first letter of struct in the reciever
// receiver type
func (o *order) changeStatus(status string) {
	o.status = status
}

func (o *order) getAmount() float32 {
	return o.amount
}

func main() {

	myOrder := order{
		id:     "1",
		amount: 100,
		status: "received",
	}
	myOrder.changeStatus("confirmed")
	fmt.Println("order struct", myOrder)
	fmt.Println("Amount:", myOrder.getAmount())
}
// use * if you want to modify the value of the struct .
// if you want to read the value of the struct then you can use value receiver.