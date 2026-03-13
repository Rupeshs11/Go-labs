package main

import (
	"fmt"
	"time"
)

type customer struct {
	name  string
	phone string
}

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
	customer
}

func main() {
	// newCustomer :=customer{
	// 	name : "john",
	// 	phone: "123456789",
	// }
	newOrder := order{
		id:     "1",
		status: "received",
		amount: 25.0,
		customer : customer{
		name : "john",
		phone: "123456789",
		},
	}

	// to change the name 
	newOrder.customer.name="knox"

	fmt.Println(newOrder)
	fmt.Println(newOrder.customer)
}
