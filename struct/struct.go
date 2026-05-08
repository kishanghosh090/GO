package main

import (
	"fmt"
	"time"
)

// order struct

type customer struct {
	name  string
	phone int
}

type order struct {
	id        int
	name      string
	amount    float32
	status    string
	createdAt time.Time // nanosecond precision
	customer
}

// reciver type
func (o *order) changeStatus(status string) {
	o.status = status
}

// constructor
func newOrder(id int, amount float32, status string) *order {
	// initial setup
	order := order{
		id:     id,
		amount: amount,
		status: status,
	}
	return &order
}

func main() {
	// var o1 order =
	// order := order{
	// 	id:        1,
	// 	name:      "chai",
	// 	amount:    20,
	// 	status:    "PENDING",
	// 	createdAt: time.Now(),
	// }

	// fmt.Println(order, order.createdAt.UTC().Local())

	// one time use structure
	language := struct {
		name   string
		isGood bool
	}{"golang", true}

	fmt.Println(language)

	// embadding
	newOrder := order{
		id:     1,
		amount: 23,
		status: "CONFIRM",
		customer: customer{
			name:  "kishan",
			phone: 9635859574,
		},
	}
	fmt.Println(newOrder)

}
