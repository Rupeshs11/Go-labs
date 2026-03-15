package main

import "fmt"

//enums are nothing but enumerated types
// there are no inbuilt enum in go
//we can implement using const

type OrderStatus string

const(
	//Received OrderStatus =iota
	Received OrderStatus="received"
	Confirmed="confirmed"
	Prepared="prepared"
	Delivered="delivered"
	//iota is a predeclared identifier representing the untyped integer ordinal number of the current const specification in a (usually parenthesized) const declaration. It is zero-indexed.

)

func changeOrderStatus(status OrderStatus){
	fmt.Println("changing order status to",status)
}

func main(){
	changeOrderStatus(Delivered)
}