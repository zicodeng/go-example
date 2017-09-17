package main

import (
	"log"
	"net/http"
	"os"
)

// MethodMux sends the request to the function
// associated with the HTTP request method.
type MethodMux struct {
	// Use a map where the key is a string (method name)
	// and the value is the associated handler function.
	HandlerFuncs map[string]func(http.ResponseWriter, *http.Request)
}

// ServeHTTP is a method of MethodMux.
// Because any struct that implements ServeHTTP method is an http.Handler,
// now our MethodMux becomes a Handler.
func (mm *MethodMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// r.Method will be the method used in the request (GET, PUT, PATCH, POST, etc.)
	fn := mm.HandlerFuncs[r.Method]
	if fn != nil {
		fn(w, r)
	} else {
		http.Error(w, "This HTTP method is not allowed", http.StatusMethodNotAllowed)
	}
}

// NewMethodMux constructs a new MethodMux
// and returns a pointer to it.
func NewMethodMux() *MethodMux {
	return &MethodMux{
		HandlerFuncs: map[string]func(http.ResponseWriter, *http.Request){},
	}
}

// HelloHandler handles requests from "/hello" resource.
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!\n"))
}

func main() {
	// Get the value of ADDR environment variable.
	addr := os.Getenv("ADDR")

	// If addr returns an empty string, default it to ":80".
	if len(addr) == 0 {
		addr = ":80"
	}

	// Create a new mux (router).
	// It is a http.Handler because it also implements ServeHTTP() method.
	// The mux calls different functions for different resource paths.
	mux := http.NewServeMux()

	// mux has three methods: HandleFunc(), Handle(), and Handler()
	// All of them can be used to handle HTTP request.
	// HandleFunc() use case:
	// Tell the mux to call the HelloHandler() function
	// when someone requests the resource path "/hello"
	// mux.HandleFunc("/hello", HelloHandler)

	methodMux := NewMethodMux()
	// HelloHandler will only be called if the request is GET.
	methodMux.HandlerFuncs["GET"] = HelloHandler

	// Handle() use case:
	// mux is a http.Handler, methodMux is also a http.Handler.
	// A handler can delegate the request to another handler.
	mux.Handle("/hello", methodMux)

	// Start the web server using the mux as router,
	// and report any errors that may occur.
	// The ListenAndServe() fucntion will block,
	// so this program will continue to run until killed.
	log.Printf("The server is listening at %s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
