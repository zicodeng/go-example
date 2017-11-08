package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type BarHandler struct {
	serviceAddr string
}

func (barHandler *BarHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprintf(w, "Bar from %s", barHandler.serviceAddr)
}

func main() {
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = ":80"
	}

	http.Handle("/", &BarHandler{addr})
	log.Printf("server is listening at http://%s...", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
