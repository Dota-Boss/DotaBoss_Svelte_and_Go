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
	http.HandleFunc("/api/scout", handler)
	log.Println("API running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
