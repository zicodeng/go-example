package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

func main() {
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = ":80"
	}

	// Read environment variables for virtual host Docker container.
	fooAddr := os.Getenv("FOO_ADDR")
	barAddr := os.Getenv("BAR_ADDR")

	mux := http.NewServeMux()
	// Set root path "/" as our gateway's only path to ensure all traffic going to reverse proxy.
	mux.Handle("/", virtualHostProxy(fooAddr, barAddr))

	log.Printf("Server is listening at https://%s...", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}

// Use reverse proxy to forward traffic to our virtual host Docker containers.
func virtualHostProxy(fooAddr, barAddr string) *httputil.ReverseProxy {
	return &httputil.ReverseProxy{
		Director: func(r *http.Request) {
			// Dynamically route traffic to a specific container
			// based on the request hostname.
			switch r.Host {
			case "foo.zicodeng.me":
				r.URL.Host = fooAddr

			case "bar.zicodeng.me":
				r.URL.Host = barAddr
			}
			r.URL.Scheme = "http"
		},
	}
}
