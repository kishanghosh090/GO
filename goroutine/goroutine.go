// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func task(id int, w *sync.WaitGroup) {
// 	defer w.Done()
// 	fmt.Println("doing task", id)
// }

// func main() {
// 	var wg sync.WaitGroup

// 	for i := 0; i <= 100; i++ {
// 		wg.Add(1)
// 		go task(i, &wg)
// 	}
// 	// time.Sleep(time.Second * 2)

// 	// wait groups goroutine
// 	wg.Wait()

// }

package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals []string

func main() {
	// goroutine
	// go greeter("hello")
	// go greeter("world")
	// time.Sleep(time.Second * 1)

	// wait groups

	websiteList := []string{
		"https://kishanranaghosh.xyz",
		"https://kishanranaghosh.xyz",
		"https://kishanranaghosh.xyz",
		"https://kishanranaghosh.xyz",
		"https://kishanranaghosh.xyz",
		"https://kishanranaghosh.xyz",
		"https://kishanranaghosh.xyz",
		"https://kishanranaghosh.xyz",
		"https://kishanranaghosh.xyz",
		"https://kishanranaghosh.xyz",
		"https://kishanranaghosh.xyz",
		"https://github.com",
		"https://google.com",
		"https://chaicode.com",
	}

	var wg sync.WaitGroup
	for _, web := range websiteList {
		wg.Add(1)
		go getSatusCode(web, &wg)
	}
	wg.Wait()
	fmt.Println(signals)
}
func getSatusCode(endpoint string, wg *sync.WaitGroup) int {
	res, err := http.Get(endpoint)
	defer wg.Done()
	if err != nil {
		panic(err)
	} else {
		signals = append(signals, endpoint)
	}
	fmt.Println("endpoint: ", endpoint, ",\nstatus: ", res.StatusCode)
	return res.StatusCode
}
func greeter(s string) {
	for i := 0; i < 100; i++ {
		fmt.Println(s)
	}
}
