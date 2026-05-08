package main

import (
	"fmt"
	"time"
)

func main() {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Printf("its weekend")
	default:
		fmt.Printf("week day")
	}

	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("integer")
			fmt.Println(t)
		}
	}

	whoAmI(true)
}
