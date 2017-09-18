package main

import (
	"log"
	"net/http"
	"os"
)

// CORSHandler is a middleware handler that wraps another http.Handler
// to do some pre- and/or post-processing of the request.
type CORSHandler struct {
	Handler http.Handler
}

// ServeHTTP is a method of CORSHandler.
// This handler handles CORS requests.
func (ch *CORSHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Use the "Origin", "Access-Control-Request-Method", or "Access-Control-Request-Headers"
	// to determine if this request should be allowed.
	// For example: restrict requests from a certain origin.
	if r.Header.Get("Origin") == "http://evil.com" {
		http.Error(w, "Sorry, requests from this origin are not allowed", http.StatusUnauthorized)
		return
	}

	// Set the various CORS response headers depending on
	// what we want our server to allow.
	w.Header().Add("Access-Control-Allow-Origin", "*")
	// ...more CORS response headers...

	// Preflight request has method "OPTIONS".
	// If this is not a preflight request,
	// just call our real handler.
	if r.Method != "OPTIONS" {
		ch.Handler.ServeHTTP(w, r)
	}
}

// NewCORSHandler wraps another handler into CORSHandler.
func NewCORSHandler(handlerToWrap http.Handler) *CORSHandler {
	return &CORSHandler{
		Handler: handlerToWrap,
	}
}

// MethodMux sends the request to the function
// associated with the HTTP request method.
type MethodMux struct {
	// Use a map where the key is a string (method name)
	// and the value is the associated handler function.
	HandlerFuncs map[string]func(w http.ResponseWriter, r *http.Request)
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

	// Wraps mux into CORSHandler
	ch := NewCORSHandler(mux)

	// Start the web server using the mux as router,
	// and report any errors that may occur.
	// The ListenAndServe() fucntion will block,
	// so this program will continue to run until killed.
	log.Printf("The server is listening at %s", addr)
	log.Fatal(http.ListenAndServe(addr, ch))
}
