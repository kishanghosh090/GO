package main

import (
	"fmt"
	"time"
)

// import (
// 	"math/rand"
// 	"time"
// )

// func processNum(numChan chan int) {

// 	for num := range numChan {
// 		println("processing number", num)
// 		time.Sleep(time.Second * 1)
// 	}
// }

// func main() {
// 	// message := make(chan string)

// 	// message <- "ping" // blocking

// 	// msg := <-message // recive

// 	// println(msg)

// 	numChan := make(chan int)

// 	go processNum(numChan)

// 	// numChan <- 5

// 	for {
// 		numChan <- rand.Intn(100)
// 	}

// }

// func sum(result chan int, num1 int, num2 int) {
// 	numResult := num1 + num2

// 	result <- numResult
// }

// func main() {

// 	result := make(chan int)

// 	go sum(result, 4, 5)

// 	res := <-result
// 	println(res)

// 	time.Sleep(time.Second * 5)

// }

// un buffer channel
// func task(done chan bool) {
// 	defer func() { done <- true }()
// 	fmt.Println("process...")
// 	time.Sleep(time.Second * 5)

// }

// func main() {
// 	done := make(chan bool)

// 	go task(done)

// 	<-done // block

// 	fmt.Println("endd....")
// }

func emailSender(emailChan chan string, done chan bool) {
	defer func() { done <- true }()
	for email := range emailChan {
		fmt.Println("sending email to... ", email)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	// email := make(chan string, 100)
	// done := make(chan bool)

	// email <- "kishan1@gmail.com"
	// email <- "kishan2@gmail.com"
	// email <- "kishan3@gmail.com"
	// email <- "kishan4@gmail.com"

	// fmt.Println(<-email)
	// fmt.Println(<-email)
	// fmt.Println(<-email)
	// fmt.Println(<-email)

	// go emailSender(email, done)

	// for i := 0; i < 10; i++ {
	// 	email <- fmt.Sprintf("%d@gmail.com", i)
	// }

	// fmt.Println("done sending")
	// close(email)
	// <-done

	///
	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10
	}()

	go func() {
		chan2 <- "GO >>> NODEJS"
	}()

	for i := 0; i < 2; i++ {
		select {
		case chan1Val := <-chan1:
			fmt.Println("recived data from chan1", chan1Val)
		case chan2Val := <-chan2:
			fmt.Println("recived data from chan1", chan2Val)
		}
	}
}
