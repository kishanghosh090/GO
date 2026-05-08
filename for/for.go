package main

import "fmt"

// for -> only construct in go for looping
func main() {
	i := 1

	for i <= 10 {
		i++
	}
	// infinite loop

	for {
		break
	}
	for i = 0; i < 11; i++ {
		fmt.Println(i + 1)
	}

	if i == 1 {
		fmt.Println("i = 1")
	}

	// 1.22 range

	for i := range 3 {
		fmt.Println(i)
	}

	var role = "admin"
	var hasPermission = true

	if role == "admin" || hasPermission {
		fmt.Print("hello it's admin")
	}

	switch i {
	case 10:
		fmt.Printf("hello %d", i)
	case 20:
	default:

	}

}
