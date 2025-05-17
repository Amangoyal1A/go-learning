package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// This is the structure you can modify to shape your custom response.
type WrappedResponse struct {
	Status string      `json:"status"`
	Source string      `json:"source"`
	Data   interface{} `json:"data"` // Change this to your custom struct if needed
}

func main() {
	http.HandleFunc("/my-api", handler)
	fmt.Println("🚀 Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// 1. Make a GET request to GitHub API
	resp, err := http.Get("https://api.github.com/repos/rails/rails")
	if err != nil {
		http.Error(w, "Failed to fetch GitHub data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// 2. Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response", http.StatusInternalServerError)
		return
	}

	// 3. Parse the JSON response (unmarshal to generic map or custom struct)
	var githubData map[string]interface{} // <-- Use a typed struct later if needed
	err = json.Unmarshal(body, &githubData)
	if err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusInternalServerError)
		return
	}

	// 4. Manipulate the data (custom transformation)
	// Here's a sample transformation: only take some fields
	customData := map[string]interface{}{
		"full_name": githubData["full_name"],
		"stars":     githubData["stargazers_count"],
		"language":  githubData["language"],
		"owner":     githubData["owner"].(map[string]interface{})["login"],
	}

	// 5. Wrap the custom response
	response := WrappedResponse{
		Status: "success",
		Source: "GitHub",
		Data:   customData, // replace this with your full struct later
	}

	// 6. Send it back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
