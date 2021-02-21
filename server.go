package main

import (
	"os"
	"runtime"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
	"encoding/json"
)

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
			Price int `json:"price"`
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
	log.Info("1 BTC = ", coin.Result.Price, " USD")
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
	
	// If APP_PORT env variable is not present
	// defaut to 8000/TCP.
	port := "8000"
	portFromEnvVar := os.Getenv("APP_PORT")
	if len(portFromEnvVar) > 0 {
		port = portFromEnvVar
	}

	// Spin up HTTP listener
	srv := &http.Server{
		Handler:      r,
		Addr:         ":" + port,
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}
	log.Info("Serving HTTP on port ", port)
	log.Fatal(srv.ListenAndServe())
}
