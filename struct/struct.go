package main

import (
	"fmt"
	"time"
)

// order struct

type order struct {
	id        int
	name      string
	amount    float32
	status    string
	createdAt time.Time // nanosecond precision
}

func main() {
	// var o1 order =
	order := order{
		id:        1,
		name:      "chai",
		amount:    20,
		status:    "PENDING",
		createdAt: time.Now(),
	}

	fmt.Println(order, order.createdAt.UTC().Local())
}
