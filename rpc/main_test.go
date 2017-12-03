package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/rpc"
	"testing"
)

// Run these benchmarks using:
// go test -bench=. -benchmem

const urlToSummarize = "http://example.com"

func BenchmarkRPC(b *testing.B) {
	// Implement a benchmark for the RPC server.
	client, err := rpc.Dial("tcp", rpcAddr)
	if err != nil {
		b.Fatalf("error dialing RPC server: %v", err)
	}
	defer client.Close()
	psum := &PageSummary{}
	for i := 0; i < b.N; i++ {
		if err := client.Call("SummaryService.GetPageSummary", urlToSummarize, psum); err != nil {
			b.Fatalf("error calling RPC: %v", err)
		}
		if psum.URL != urlToSummarize {
			b.Fatalf("incorrect data returned from RPC, expected %s but got %s", urlToSummarize, psum.URL)
		}
	}
}

func BenchmarkHTTP(b *testing.B) {
	// Implement a benchmark for the HTTP server.
	summaryURL := fmt.Sprintf("http://%s?url=%s", httpAddr, urlToSummarize)
	psum := &PageSummary{}
	// N value will be automatically adjusted by Go testing library.
	// It will keep running until get a consistent result.
	for i := 0; i < b.N; i++ {
		resp, err := http.Get(summaryURL)
		if err != nil {
			b.Fatalf("error getting page summary: %v", err)
		}
		if err := json.NewDecoder(resp.Body).Decode(psum); err != nil {
			b.Fatalf("error decoding JSON response: %v", err)
		}
		if psum.URL != urlToSummarize {
			b.Fatalf("incorrect URL returned: expected %s but got %s", urlToSummarize, psum.URL)
		}
		resp.Body.Close()
	}
}
