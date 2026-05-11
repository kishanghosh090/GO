package main

import (
	"io"
	"net/http"
)

const url = "https://kishanranaghosh.xyz"

func main() {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	// var b []byte
	databyte, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	println(string(databyte))
}
