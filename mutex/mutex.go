package main

import (
	"fmt"
	"sync"
)

// solve race condi using mutex ---------->>>>>>>>>>>>>>>>>>>>>>>>>
type Post struct {
	id    int
	views int

	mu sync.Mutex
}

func (p *Post) inc(wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	}()
	p.mu.Lock()
	p.views++
	p.mu.Unlock()
}

func main() {
	myPost := Post{views: 0}
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go myPost.inc(&wg)
	}
	wg.Wait()

	// myPost.inc()
	// myPost.inc()
	fmt.Print(myPost.views)

	// time.Sleep(time.Second * 3)
}
