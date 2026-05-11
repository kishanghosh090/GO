package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("")
	PerformGetReq()
}

func PerformGetReq() {
	const myUrl = "http://localhost:4002/get"

	response, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println(strings.Split(response.Status, " ")[0] == "200")

	content, _ := io.ReadAll(response.Body)

	println(string(content))

}
