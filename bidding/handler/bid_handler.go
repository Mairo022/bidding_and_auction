package handler

import (
	"encoding/json"
	"math/rand"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type AdObject struct {
	AdID     uuid.UUID `json:"ad_id"`
	BidPrice int       `json:"bid_price"`
}

func BidHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("received request")
	reject := rand.Intn(2)

	if reject == 1 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	bidPrice := rand.Intn(10000)
	adObject := AdObject{
		AdID:     uuid.New(),
		BidPrice: bidPrice,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(adObject)
}
