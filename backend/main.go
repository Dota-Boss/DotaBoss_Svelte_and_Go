package main

import (
	"encoding/json"
	"fmt"
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

func scoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ScoutRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	fmt.Printf("Received scout request for: %s vs %s\n", req.TeamA, req.TeamB)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	resp := ScoutResponse{
		MatchupRating: "7.8/10",
		SuggestedBan:  "Invoker",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/api/scout", scoutHandler)
	fmt.Println("Backend running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
