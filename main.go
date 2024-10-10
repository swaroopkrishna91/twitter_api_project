package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"
)

// Function to delete a tweet by ID
func deleteTweet(client *http.Client, tweetID string) {
	// Set the URL for deleting the tweet
	url := "https://api.twitter.com/2/tweets/" + tweetID

	// Create the request
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	// Set the required headers
	req.Header.Set("Authorization", "Bearer "+os.Getenv("TWITTER_BEARER_TOKEN"))
	req.Header.Set("Content-Type", "application/json")

	// Make the DELETE request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to delete tweet: %v", err)
	}
	defer resp.Body.Close()

	// Check the response
	if resp.StatusCode != http.StatusNoContent {
		log.Fatalf("New tweet posted and deleted: %v\n", resp.Status)
	}

	// Successfully deleted the tweet
	fmt.Printf("Tweet with ID %s deleted.\n", tweetID)
}

func main() {

	//Loading .env File
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Twitter API credentials
	consumerKey := os.Getenv("CONSUMER_KEY")
	consumerSecret := os.Getenv("CONSUMER_SECRET")
	accessToken := os.Getenv("ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("ACCESS_TOKEN_SECRET")

	// OAuth1 authentication setup
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessTokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	// Post a tweet
	// Define the tweet payload
	tweetstring := "Personal Tweet Test"
	payload := map[string]interface{}{
		"text": tweetstring,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("Error marshaling payload: %v", err)
	}

	// Create the request to post a tweet
	url := "https://api.twitter.com/2/tweets"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	// Set the required headers
	req.Header.Set("Content-Type", "application/json")

	// Make the request
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatalf("Failed to post tweet: %v", err)
	}
	defer resp.Body.Close()

	// Check the response
	if resp.StatusCode != http.StatusCreated {
		log.Fatalf("Failed to post tweet status: %s", resp.Status)
	}

	// Read and print the response
	var result map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Fatalf("Error decoding response: %v", err)
	}
	tweetID := result["data"].(map[string]interface{})["id"].(string)

	fmt.Printf("Tweet ID: %v\n", result["data"].(map[string]interface{})["id"])

	//  Sleep for 30 seconds
	time.Sleep(30 * time.Second)

	// Delete the tweet
	deleteTweet(httpClient, tweetID)
}
