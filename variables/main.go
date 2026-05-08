package main

import "fmt"

const age = 20

func main() {
	fmt.Println("variables")
	var user string = "hello from kishan"
	const name string = "hello from chai code"

	var age = 19
	if age != 0 {

	}

	myName := "kishan"
	fmt.Println(myName)
	for i := 0; i < 11; i++ {
		fmt.Println(i)
	}
	var isLoggedIn bool = true // type inference, the type of isLoggedIn will be bool
	fmt.Println(user)
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of type: %T \n", user)       // %T is used to print the type of the variable
	fmt.Printf("variable is of type: %T \n", isLoggedIn) // %T is used to print the type of the variable

	var smallValue uint8 = 255                                                       // uint8 can hold values from 0 to 255
	fmt.Printf("smallValue is of type: %T and value: %d \n", smallValue, smallValue) // %d is used to print the value of the variable

	var largeValue uint64 = 18446744073709551615                                     // uint64 can hold values from 0 to 18446744073709551615
	fmt.Printf("largeValue is of type: %T and value: %d \n", largeValue, largeValue) // %d is used to print the value of the variable

	var myRoll int
	myRoll = 12

	fmt.Println(myRoll)

	// new syntax
	const (
		port = 5000
		host = "localhost"
	)
	fmt.Println(port)
}
