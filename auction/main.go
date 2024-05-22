package main

import (
	"bidding_and_auction/auction/handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/auction", handler.AuctionHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
