package services

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

type ApiService interface {
	GetApi() (Post, error)
}

type apiService struct{}

func NewApiService() ApiService {
	return &apiService{}
}

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func (s *apiService) GetApi() (Post, error) {
	// Create a Resty client
	client := resty.New()

	// Make a GET request to JSONPlaceholder
	resp, err := client.R().
		SetResult(&Post{}).
		Get("https://jsonplaceholder.typicode.com/posts/1")

	if err != nil {
		log.Fatal("API call failed:", err)
	}
	// Print original data
	fmt.Println("🔹 started Response:",resp)

	// Parse response into Post struct
	var post Post
	if err := json.Unmarshal(resp.Body(), &post); err != nil {
		return Post{}, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	// Print original data
	fmt.Println("🔹 Original Response:")
	fmt.Printf("%+v\n\n", post)

	// Modify data
	post.Title = "🚀 Modified Title"
	post.Body = "This body was manipulated in Go."

	// Print modified data
	fmt.Println("🛠️ Modified Response:")
	fmt.Printf("Title: %s\n", post.Title)
	fmt.Printf("Body: %s\n", post.Body)

	return post, nil
}
