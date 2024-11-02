package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type PRData struct {
	URL string `json:"url"`
}

type File struct {
	SHA         string `json:"sha"`
	Filename    string `json:"filename"`
	Status      string `json:"status"`
	Additions   int    `json:"additions"`
	Deletions   int    `json:"deletions"`
	Changes     int    `json:"changes"`
	BlobURL     string `json:"blob_url"`
	RawURL      string `json:"raw_url"`
	ContentsURL string `json:"contents_url"`
	Patch       string `json:"patch"`
}

type Response struct {
	Files []File `json:"files"`
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

	fmt.Printf("Repository URL: %s\n", prData.URL)

	var bearer = "Bearer " + os.Getenv("GITHUB_ACCESS_TOKEN")

	req, _ := http.NewRequest("GET", prData.URL, nil)
	req.Header.Add("Authorization", bearer)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error on response. ", err)
	}
	defer resp.Body.Close()

	var files []File
	if err := json.NewDecoder(resp.Body).Decode(&files); err != nil {
		fmt.Println("error decoding response: ", err)
	}

	fmt.Printf("the files are %+v", files)
}
