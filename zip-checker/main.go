package main

import (
	"fmt"
	"github.com/zicodeng/go-example/zip-checker/handlers"
	"github.com/zicodeng/go-example/zip-checker/models"
	"log"
	"net/http"
	"os"
	"strings"
)

const zipsPath = "/zips/"

func main() {
	// Reading ADDR environment variable from OS.
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = ":443"
	}

	tlskey := os.Getenv("TLSKEY")
	tlscert := os.Getenv("TLSCERT")
	if len(tlskey) == 0 || len(tlscert) == 0 {
		log.Fatal("Please set TLSKEY and TLSCERT")
	}

	zips, err := models.LoadZips("./client/zips.csv")
	if err != nil {
		log.Fatalf("error loading zips: %v", err)
	}

	log.Printf("Loaded %d zips", len(zips))

	// Index all zips by city.
	cityIndex := models.ZipIndex{}
	for _, z := range zips {
		cityLower := strings.ToLower(z.City)
		cityIndex[cityLower] = append(cityIndex[cityLower], z)
	}

	mux := http.NewServeMux()

	cityHandler := &handlers.CityHandler{
		PathPrefix: zipsPath,
		Index:      cityIndex,
	}

	mux.Handle("/", http.FileServer(http.Dir("/client")))
	mux.Handle(zipsPath, cityHandler)

	fmt.Printf("Server is listening at https://%s\n", addr)
	log.Fatal(http.ListenAndServeTLS(addr, tlscert, tlskey, mux))
}
