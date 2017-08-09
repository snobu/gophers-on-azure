package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

var log = logrus.New()

func statsHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("In statsHandler..")
	stats := "some stats as string"
	fmt.Fprintf(w, "%s", stats)
}

func main() {
	var dir string
	flag.StringVar(&dir, "dir", ".", "Serve from current directory")
	flag.Parse()
	r := mux.NewRouter()
	r.HandleFunc("/stats", statsHandler).Methods("GET")
	r.PathPrefix("/").
		Handler(http.FileServer(http.Dir(dir)))
	srv := &http.Server{
		Handler:      r,
		Addr:         "localhost:8000",
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}
	log.Info("Serving HTTP on port 8000.")
	log.Fatal(srv.ListenAndServe())
}
