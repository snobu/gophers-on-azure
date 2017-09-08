package main

import (
	"runtime"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
	"encoding/json"
)

var log = logrus.New()

func statsHandler(w http.ResponseWriter, r *http.Request) {
	stats := runtime.GOOS + "\n" +
			 runtime.GOARCH + "\n" +
			 runtime.Version()
	fmt.Fprintf(w, "%s", stats)
}

func priceHandler(w http.ResponseWriter, r *http.Request) {
    url := "https://api.cryptowat.ch/markets/kraken/btcusd/price"

	type Coin struct {
		Result struct {
			Price float64 `json:"price"`
		} `json:"result"`
	}

	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	var coin Coin

	if err := json.NewDecoder(resp.Body).Decode(&coin); err != nil {
		log.Println(err)
	}

	fmt.Fprintf(w, "%.2f", coin.Result.Price)

}

func main() {
	var dir string
	flag.StringVar(&dir, "dir", ".", "Serve from current directory")
	flag.Parse()
	r := mux.NewRouter()
	r.HandleFunc("/stats", statsHandler).Methods("GET")
	r.HandleFunc("/price", priceHandler).Methods("GET")
	r.PathPrefix("/").
		Handler(http.FileServer(http.Dir(dir)))
	srv := &http.Server{
		Handler:      r,
		Addr:         ":8000",
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}
	log.Info("Serving HTTP on port 8000.")
	log.Fatal(srv.ListenAndServe())
}
