package handlers

import (
	"fmt"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if len(name) == 0 {
		http.Error(w, "no query found in the requested URL", http.StatusBadRequest)
	}

	w.Header().Add("Access-Control-Allow-Origin", "*")

	fmt.Fprintf(w, "Hello, %s!", name)
}
