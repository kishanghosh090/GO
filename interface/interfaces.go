package main

import "fmt"

type Payment interface {
	pay(amount float32)
	refund(amount float32)
}

type payment struct {
	gateway Payment
}

// razorpay
type razorpay struct{}

func (r razorpay) pay(amount float32) {
	fmt.Println("making payment using raz", amount)
}
func (r razorpay) refund(amount float32) {
	fmt.Println("making payment using raz", amount)
}

// paypal
type paypal struct{}

func (r paypal) pay(amount float32) {
	fmt.Println("making payment using raz", amount)
}

func main() {
	raz := razorpay{}

	pay := payment{
		gateway: raz,
	}
	pay.gateway.pay(1000)
}
