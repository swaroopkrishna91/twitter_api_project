package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

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
		log.Fatalf("A new tweet has been created and deleted as well: %v\n", resp.Status)
	}

	// Successfully deleted the tweet
	fmt.Printf("Tweet with ID %s deleted successfully.\n", tweetID)
}

func main() {
	// Twitter API credentials

	//Loading .env File
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	consumerKey := os.Getenv("CONSUMER_KEY")
	consumerSecret := os.Getenv("CONSUMER_SECRET")
	accessToken := os.Getenv("ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("ACCESS_TOKEN_SECRET")

	// OAuth1 authentication setup
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessTokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	// client := twitter.NewClient(httpClient)

	// Post a tweet

	// tweet, _, err := client.Statuses.Update("Hello from Twitter API!", nil)
	// if err != nil {
	// log.Fatalf("Failed to post tweet: %v", err)
	// }
	// fmt.Printf("Successfully posted tweet with ID: %d\n", tweet.ID)

	// tweet, _, err := client.Statuses.Update("Hello, Twitter! This is my first tweet using Go!", nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Tweet posted: %s\n", tweet.Text)

	// Define the tweet payload
	tweetText := "Hello Twitter API V2! This is an EXTERNAL tweet from WINP2000 Team."
	payload := map[string]interface{}{
		"text": tweetText,
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

	fmt.Printf("Posted Tweet ID: %v\n", result["data"].(map[string]interface{})["id"])

	// Delete the tweet (optional)
	deleteTweet(httpClient, tweetID)
}
