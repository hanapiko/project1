package handlers

import (
	"encoding/json"
	"group/utils"
	"log"
	"net/http"
)

// DateEntry represents a single date entry with its ID and dates
type DateEntry struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

// DateResponse represents the response structure for the /dates endpoint
type DateResponse struct {
	Index []DateEntry `json:"index"`
}

// DatesHandler handles requests to the /dates endpoint
func DatesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Fetch data from the external API
	data, err := utils.FetchDataFromAPI("https://groupietrackers.herokuapp.com/api/dates")
	if err != nil {
		log.Printf("Error fetching data: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Parse the JSON data
	var dateResponse DateResponse
	if err := json.Unmarshal(data, &dateResponse); err != nil {
		log.Printf("Error parsing JSON: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Respond with JSON data
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(dateResponse); err != nil {
		log.Printf("Error encoding JSON: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
