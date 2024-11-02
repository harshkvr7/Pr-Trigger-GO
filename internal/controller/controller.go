package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type PullRequest struct {
	Number int `json:"number"`
}

type Repository struct {
	URL string `json:"url"`
}

type Event struct {
	PullRequest PullRequest `json:"pull_request"`
	Repository  Repository  `json:"repository"`
}

type ChangedFile struct {
	Filename string `json:"filename"`
}

func SendGreeting(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Greetings from server.")
}

func GetPrDetails(w http.ResponseWriter, r *http.Request) {
	var rawBody []byte
	rawBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	fmt.Println("Raw Event Body:", string(rawBody))

	// var event Event
	// if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
	// 	http.Error(w, "Invalid request payload", http.StatusBadRequest)
	// 	return
	// }

	// prNumber := event.PullRequest.Number
	// filesUrl := fmt.Sprintf("%s/pulls/%d/files", event.Repository.URL, prNumber)

	// fmt.Println("url :", filesUrl)

	// req, err := http.NewRequest("GET", filesUrl, nil)
	// if err != nil {
	// 	http.Error(w, "Failed to create request", http.StatusInternalServerError)
	// 	return
	// }

	// fmt.Println("ACCESS TOKEN:", os.Getenv("GITHUB_ACCESS_TOKEN"))

	// req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("GITHUB_ACCESS_TOKEN")))
	// req.Header.Set("Accept", "application/vnd.github.v3+json")

	// client := &http.Client{}
	// resp, err := client.Do(req)
	// if err != nil {
	// 	http.Error(w, "Failed to fetch files 1", http.StatusInternalServerError)
	// 	println(err)
	// 	return
	// }
	// defer resp.Body.Close()

	// if resp.StatusCode != http.StatusOK {
	// 	http.Error(w, "Failed to fetch files 2", http.StatusInternalServerError)
	// 	return
	// }

	// var changedFiles []ChangedFile
	// if err := json.NewDecoder(resp.Body).Decode(&changedFiles); err != nil {
	// 	http.Error(w, "Failed to decode response", http.StatusInternalServerError)
	// 	return
	// }

	// fileNames := make([]string, len(changedFiles))
	// for i, file := range changedFiles {
	// 	fileNames[i] = file.Filename
	// }

	// fmt.Println("changed files :", fileNames)

	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode("done")
}
