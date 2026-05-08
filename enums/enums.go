package main

// type MyType string
type Status int

const (
	RECIVED Status = iota
	CONFIRMED
	PREPARED
	DELIVERED
)

func changeOrderStatus(status Status) {}

func main() {
	changeOrderStatus(RECIVED)
}
