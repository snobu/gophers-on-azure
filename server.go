package main

import (
	"flag"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

var router *mux.Router
var routes []string

func main() {
	var dir string
	flag.StringVar(&dir, "dir", ".", "Serve from current directory")
	flag.Parse()
	router = mux.NewRouter()
	router.HandleFunc("/env", envHandler).Methods("GET")
	router.HandleFunc("/price", priceHandler).Methods("GET")
	router.HandleFunc("/routes", gorillaWalkHandler).Methods("GET")
	router.PathPrefix("/").Methods("GET").Handler(http.FileServer(http.Dir(dir)))

	// If APP_PORT env variable is not present
	// defaut to 8000/TCP.
	port := "8000"
	portFromEnvVar := os.Getenv("APP_PORT")
	if len(portFromEnvVar) > 0 {
		port = portFromEnvVar
	}

	// Spin up HTTP listener
	srv := &http.Server{
		Handler:      router,
		Addr:         ":" + port,
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}
	log.Info("Serving HTTP on port ", port)
	log.Fatal(srv.ListenAndServe())
}