package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

const rpcAddr = "localhost:6000"
const httpAddr = "localhost:4000"

// PreviewImage represents a page summary preview image.
type PreviewImage struct {
	URL    string
	Alt    string
	Width  int
	Height int
}

// PageSummary represents a summary of a web page.
type PageSummary struct {
	URL         string
	Title       string
	Description string
	Previews    []*PreviewImage
}

// newTestSummary returns a new PageSummary for testing purposes.
func newTestSummary(pageURL string) *PageSummary {
	return &PageSummary{
		URL:         pageURL,
		Title:       "Test title",
		Description: "Test description",
		Previews: []*PreviewImage{
			{
				URL:    pageURL + "/test.png",
				Alt:    "A test image",
				Width:  10,
				Height: 20,
			},
		},
	}
}

// Implement an RPC service
// and an HTTP handler function that
// both accept a URL and return a PageSummary.

func SummaryHandler(w http.ResponseWriter, r *http.Request) {
	pageSummary := newTestSummary(r.FormValue("url"))
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pageSummary)
}

type SummaryService struct{}

// A valid RPC service method can only accept two arguments:
// 1. arg type
// 2. reply type
func (ss *SummaryService) GetPageSummary(pageURL string, pageSummary *PageSummary) error {
	// Copy over the value.
	*pageSummary = *newTestSummary(pageURL)
	return nil
}

func startRPC(addr string) {
	// Start RPC server.
	svc := &SummaryService{}
	rpc.Register(svc)
	// Listen directly on TCP socket on the server side.
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("error binding to RPC port: %v", err)
	}
	log.Printf("RPC server is listening at %s", addr)
	rpc.Accept(lis)
}

func main() {
	// Start the RPC server on rpcAddr.
	go startRPC(rpcAddr)

	mux := http.NewServeMux()
	mux.HandleFunc("/", SummaryHandler)
	log.Printf("HTTP server is listening at %s\n", httpAddr)
	log.Fatal(http.ListenAndServe(httpAddr, mux))
}
