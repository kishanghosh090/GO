package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kishanghosh090/mongoAPI/router"
)

func main() {
	fmt.Println("Mongodb API")

	r := router.Router()
	fmt.Println("Server is getting started...")
	println("Listing at port 4000")
	log.Fatal(http.ListenAndServe(":4000", r))

}
