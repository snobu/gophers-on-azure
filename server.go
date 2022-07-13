package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func envHandler(w http.ResponseWriter, r *http.Request) {
	var pairs []string
	for _, e := range os.Environ() {
        pairs = append(pairs, e)
    }

	json.NewEncoder(w).Encode(pairs)
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

	fmt.Fprintf(w, "%d", int(coin.Result.Price))
	log.Info("1 BTC = ", coin.Result.Price, " USD")
}

func main() {
	var dir string
	flag.StringVar(&dir, "dir", ".", "Serve from current directory")
	flag.Parse()
	r := mux.NewRouter()
	r.HandleFunc("/env", envHandler).Methods("GET")
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
