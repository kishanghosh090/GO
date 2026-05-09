package main

import (
	"fmt"
	"time"
)

func task(id int) {
	fmt.Println("doing task", id)

}

func main() {
	for i := 0; i <= 100; i++ {
		go task(i)
	}
	time.Sleep(time.Second * 2)
}
