package main

import (
	"github.com/zicodeng/go-example/automated-testing/handlers"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", handlers.HelloHandler)

	log.Println("server is listening at http://localhost:3000")
	log.Fatal(http.ListenAndServe("localhost:3000", mux))
}
