package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func envHandler(w http.ResponseWriter, r *http.Request) {
	// Unnecessary since it will default to text/plain anyway
	w.Header().Set("Content-Type", "text/plain")
	for _, pair := range os.Environ() {
		fmt.Fprintf(w, "%s\n", pair)
	}
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

func gorillaWalkHandler(w http.ResponseWriter, r *http.Request) {
	err := router.Walk(
		func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		path, err := route.GetPathTemplate()
		if err != nil {
			log.Fatal(err)
		}
		routes = append(routes, path)
	
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	for _, route := range routes {
		fmt.Fprintf(w, "%s\n", route)
	}
	routes = []string{}
}