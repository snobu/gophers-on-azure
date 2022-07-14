package main

import (
	"encoding/json"
	"net/http"
	"os"
)

func envHandler(w http.ResponseWriter, r *http.Request) {
	var pairs []string
	for _, e := range os.Environ() {
		pairs = append(pairs, e)
	}

	json.NewEncoder(w).Encode(pairs)
}