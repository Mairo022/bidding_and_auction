package main

import (
	"bidding_and_auction/bidding/handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/bid", handler.BidHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
