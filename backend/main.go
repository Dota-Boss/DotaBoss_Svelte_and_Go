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

func withCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Allow requests from anywhere (or set to "http://localhost:5173" if you want)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")

		// Respond to preflight OPTIONS request
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// Call the actual handler
		next(w, r)
	}
}

func scoutHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

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

	resp := ScoutResponse{
		MatchupRating: "7.8/10",
		SuggestedBan:  "Invoker",
	}

	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/api/scout", withCORS(scoutHandler))
	fmt.Println("Backend running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
