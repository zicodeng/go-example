package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type ByeHandler struct {
	serviceAddr string
}

func (byeHandler *ByeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprintf(w, "Bye from %s", byeHandler.serviceAddr)
}

func main() {
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = ":80"
	}

	http.Handle("/v1/bye", &ByeHandler{addr})
	log.Printf("server is listening at http://%s...", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
