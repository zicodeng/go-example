package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type FooHandler struct {
	serviceAddr string
}

func (fooHandler *FooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprintf(w, "Foo from %s", fooHandler.serviceAddr)
}

func main() {
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = ":80"
	}

	http.Handle("/", &FooHandler{addr})
	log.Printf("server is listening at http://%s...", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
