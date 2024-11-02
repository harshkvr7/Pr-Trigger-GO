package controller

import (
	"encoding/json"
	"net/http"
)

func SendGreeting(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Greetings from server.")
}
