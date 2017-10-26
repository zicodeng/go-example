package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func CurrentTimeHandler(w http.ResponseWriter, r *http.Request) {
	curTime := time.Now().Format(time.Kitchen)
	w.Write([]byte(fmt.Sprintf("the current time is %v", curTime)))
}

func main() {
	addr := os.Getenv("ADDR")

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/hello", HelloHandler)
	mux.HandleFunc("/v1/time", CurrentTimeHandler)

	// Wrap entire mux with logger middleware.
	// This is valid because http.ServeMux satisfies the http.Handler interface.
	wrappedMux := NewLogger(mux)

	log.Printf("server is listening at %s", addr)

	// http.ListenAndServe expects a http.Handler as its second parameter.
	// Since a Logger has http.ServeHTTP method,
	// so it also satisfies the http.Handler interface.
	log.Fatal(http.ListenAndServe(addr, wrappedMux))
}

// Logger is a middleware handler that does request logging.
type Logger struct {
	handler http.Handler
}

// NewLogger constructs a new Logger middleware handler.
func NewLogger(handlerToWrap http.Handler) *Logger {
	return &Logger{handlerToWrap}
}

// ServeHTTP handles the request by passing it to the real
// handler and logging the request details.
func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Pre-process the request.
	start := time.Now()

	// Handle the request by calling the real handler.
	l.handler.ServeHTTP(w, r)

	// Post-process the request.
	log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
}
