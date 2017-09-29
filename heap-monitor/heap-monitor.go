package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime"
)

func memoryHandler(w http.ResponseWriter, r *http.Request) {
	stats := &runtime.MemStats{}
	runtime.ReadMemStats(stats)

	// Add response header.
	w.Header().Add("Content-Type", "application/json")
	// Enable CORS.
	w.Header().Add("Access-Control-Allow-Origin", "*")

	// Send the JSON to front-end.
	json.NewEncoder(w).Encode(stats)
}

func gcHandler(w http.ResponseWriter, r *http.Request) {
	// Manually force the garbage collector to run.
	runtime.GC()
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/memory", memoryHandler)
	mux.HandleFunc("/memory/gc", gcHandler)
	fmt.Printf("Server is listening at http://localhost:3000\n")
	log.Fatal(http.ListenAndServe("localhost:3000", mux))
}
