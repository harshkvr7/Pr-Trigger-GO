package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PRData struct {
	Number int    `json:"number"`
	URL    string `json:"url"`
}

func SendGreeting(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Greetings from server.")
}

func GetPrDetails(w http.ResponseWriter, r *http.Request) {
	var prData PRData

	err := json.NewDecoder(r.Body).Decode(&prData)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	fmt.Printf("PR Number: %d, Repository URL: %s\n", prData.Number, prData.URL)
}
