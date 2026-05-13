
package main

import (
	"log"
	"net/http"

	"github.com/kishanghosh090/api/internal/config"
)

func main() {
	cfg := config.MustLoad()

	// setup router

	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome to chai code go server"))
	})
	// println(cfg.Addr)
	log.Fatal(http.ListenAndServe(cfg.Addr, router))
}
