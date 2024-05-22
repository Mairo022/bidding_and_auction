package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type BiddingServiceResponse struct {
	AdID     string `json:"ad_id"`
	BidPrice int    `json:"bid_price"`
}

func AuctionHandler(w http.ResponseWriter, r *http.Request) {
	adPlacementID := r.URL.Query().Get("ad_placement_id")
	log.Println("request for item:", adPlacementID)

	bidServicePorts := [3]int{8081, 8082, 8083}
	bids := make([]BiddingServiceResponse, len(bidServicePorts))

	var wg sync.WaitGroup

	for i, port := range bidServicePorts {
		bidServiceURL := fmt.Sprintf("http://localhost:%d/bid", port)

		wg.Add(1)
		go func(url string, j int) {
			defer wg.Done()

			bidResponse, err := callBiddingService(url)
			if err != nil {
				return
			}

			bids[j] = *bidResponse
		}(bidServiceURL, i)
	}
	wg.Wait()

	var bestBid *BiddingServiceResponse

	for _, bid := range bids {
		if bid.AdID == "" {
			continue
		}
		if bestBid == nil || bid.BidPrice > bestBid.BidPrice {
			localBid := bid
			bestBid = &localBid
		}
	}


	if bestBid == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bestBid)
}

func callBiddingService(url string) (*BiddingServiceResponse, error) {
	client := &http.Client{
		Timeout: 200 * time.Millisecond,
	}

	resp, err := client.Get(url)

	if err != nil {
		fmt.Println("error making request", err)
		return nil, err
	}
	defer resp.Body.Close()

	var bidResponse BiddingServiceResponse
	if err := json.NewDecoder(resp.Body).Decode(&bidResponse); err != nil {
		log.Println("response decoding error for", url)
		return nil, err
	}

	return &bidResponse, nil
}
