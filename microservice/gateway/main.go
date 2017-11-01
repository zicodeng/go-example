package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
	"sync"
)

// NewServiceProxy returns a reverse proxy that balances load.
func NewServiceProxy(addrs []string) *httputil.ReverseProxy {
	i := 0
	mutex := sync.Mutex{}
	return &httputil.ReverseProxy{
		Director: func(r *http.Request) {
			// A Mutex lock is used here to prevent concurrent access.
			mutex.Lock()
			// For incoming requests, forward them to different microservice instances.
			r.URL.Host = addrs[i%len(addrs)]
			i++
			mutex.Unlock()
			r.URL.Scheme = "http"
		},
	}
}

// RootHandler handles requests for the root resource.
func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprintf(w, "Hello from the gateway! Try requesting /v1/hello or /v1/bye")
}

func main() {
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = ":443"
	}

	// Read environment variables.
	helloAddrs := os.Getenv("HELLO_ADDRS")
	helloAddrSlice := strings.Split(helloAddrs, ",")

	byeAddrs := os.Getenv("BYE_ADDRS")
	byeAddrSlice := strings.Split(byeAddrs, ",")

	mux := http.NewServeMux()
	mux.HandleFunc("/", RootHandler)
	// Add reverse proxy handler for `/v1/hello` and `/v1/bye`
	mux.Handle("/v1/hello", NewServiceProxy(helloAddrSlice))
	mux.Handle("/v1/bye", NewServiceProxy(byeAddrSlice))

	log.Printf("Server is listening at https://%s...", addr)
	log.Fatal(http.ListenAndServeTLS(addr, "tls/fullchain.pem", "tls/privkey.pem", mux))
}
