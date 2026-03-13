package main

import "fmt"

// interface creation

type paymenter interface{
	pay(amount float32)
	refund(amount float32,id string)
}

type payment struct {
	gateway paymenter
}

func (p payment) makePayment(amount float32) {
	//razorpayPaymentGw := razorpay{}
	//stripePaymentGw := stripe{}
	//razorpayPaymentGw.pay(amount)
	//stripePaymentGw.pay(amount)
	p.gateway.pay(amount)

}
func (p payment) refundPayment(amount float32,id string){
	p.gateway.refund(amount,id)
}

type razorpay struct {}

func (r razorpay) pay(amount float32) {
	// logic to make payment
	fmt.Println("making payment using razorpay", amount)
}

type stripe struct{}

func (s stripe) pay(amount float32){
	fmt.Println("making payment using stripe",amount)
}

type paypal struct{}

func (p paypal) pay(amount float32){
	fmt.Println("making payment using paypal for",amount)
}

type googlepay struct{}

func (g googlepay) pay(amount float32){
	fmt.Println("making payment using Googlepay for",amount)
}

func (g googlepay) refund(amount float32,id string){
	fmt.Println("initiating the refund using  googlepay of",amount,id)
}

type phonepay struct{}

func (p phonepay) pay(amount float32){
	fmt.Println("making payment using phonepay for",amount)
}

type fakepayment struct{}

func (f fakepayment) pay(amount float32){
	fmt.Println("making paymnet using fake payment gateway for testing",amount)
}

func main() {
	//stripePaymentGw := stripe{}
	//razorpayPaymentGw := razorpay{}
	//paypalPaymentGw := paypal{}
	//phonepayPaymentGw := phonepay{}
	googlepayPaymentGw := googlepay{}
	//fakeGw := fakepayment{}
	newPayment := payment{
		//gateway: razorpayPaymentGw,
		//gateway: stripePaymentGw,
		gateway: googlepayPaymentGw,
	}
	newPayment.makePayment(100)
	newPayment.refundPayment(100,"11")
}
