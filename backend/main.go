package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type ScoutRequest struct {
	TeamA string `json:"teamA"`
	TeamB string `json:"teamB"`
}

type ScoutResponse struct {
	MatchupRating string `json:"matchupRating"`
	SuggestedBan  string `json:"suggestedBan"`
}

func withCORS(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		h.ServeHTTP(w, r)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST", http.StatusMethodNotAllowed)
		return
	}

	var req ScoutRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	resp := ScoutResponse{"7.0/10", "Invoker"}
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/api/scout", withCORS(handler))
	log.Println("API running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
