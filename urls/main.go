package main

import (
	"fmt"
	"net/url"
)

const myURI = "https://localhost:4002/?id=1"

func main() {
	fmt.Println("handling URL in GoLang")
	// parsing

	result, _ := url.Parse(myURI)

	fmt.Println(result.Scheme)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)

	qparams := result.Query()

	// fmt.Println("%T", qparams)
	println(qparams["id"])

	partsURI := &url.URL{
		Scheme: "https",
		Host:   "localhost:4002",
		Path:   "/",
	}
	anotherURI := partsURI.String()
	println(anotherURI)
}
