package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race condition - kishanranaghosh.xyz")
	var score = []int{0}

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		score = append(score, 1)
		defer wg.Done()
	}(wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		score = append(score, 2)
		defer wg.Done()
	}(wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		score = append(score, 3)
		defer wg.Done()
	}(wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		score = append(score, 4)
		defer wg.Done()
	}(wg)

	wg.Wait()

}
