package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/zicodeng/go-example/zip-checker/models"
	"net/http"
	"strings"
)

// CityHandler is a http.Handler that has a required ServeHTTP method
// and additional data.
type CityHandler struct {
	PathPrefix string
	Index      models.ZipIndex
}

func (ch *CityHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// URL pattern: /zips/city-name
	cityName := r.URL.Path[len(ch.PathPrefix):]
	cityName = strings.ToLower(cityName)
	if len(cityName) == 0 {
		// log.Fatal will terminate the server, don't use it for logging http error info.
		http.Error(w, "Please provide a city name", http.StatusBadRequest)
		return
	}

	w.Header().Add(headerContentType, contentTypeJSON)
	w.Header().Add(headerAccessControlAllowOrigin, "*")

	// Get all zip codes for a given city.
	zips := ch.Index[cityName]
	if len(zips) == 0 {
		json.NewEncoder(w).Encode(fmt.Sprintf("No zip code found at %s.", cityName))
		return
	}
	json.NewEncoder(w).Encode(zips)
}
