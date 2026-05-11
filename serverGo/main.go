package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("")
	// PerformGetReq()
	// PerformPostReq()
	// PerformPostFormRequest()

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

func PerformPostReq() {
	const myUrl = "http://localhost:4002/post"

	// fake json payload

	reqBody := strings.NewReader(`
		{
			"courseName" : "lets go ",
			"price" : 0
		}
	`)

	response, err := http.Post(myUrl, "application/json", reqBody)

	if err != nil {
		panic(err)
	}

	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))

}

func PerformPostFormRequest() {
	const myUrl = "http://localhost:4002/post"

	data := url.Values{}

	data.Add("name", "elychi")
	data.Add("type", "cold chai")

	res, err := http.PostForm(myUrl, data)

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	content, _ := io.ReadAll(res.Body)
	fmt.Println(string(content))
}
