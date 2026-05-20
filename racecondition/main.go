package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race condition - kishanranaghosh.xyz")
	var score = []int{0}

	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}

	wg.Add(1)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		defer wg.Done()
	}(wg, mut)

	wg.Add(1)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		defer wg.Done()
	}(wg, mut)

	wg.Add(1)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		defer wg.Done()
	}(wg, mut)

	wg.Add(1)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		mut.Lock()
		score = append(score, 4)
		mut.Unlock()
		defer wg.Done()
	}(wg, mut)

	wg.Wait()

}
