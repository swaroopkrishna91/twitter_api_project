# Twitter Tweet Post and Delete Script

This Go application demonstrates how to post a tweet using Twitter API v2 and delete the tweet after a delay of 30 seconds.

## Requirements

Create an `.env` file in the root directory of the project (this is where you will store your Twitter API credentials).

## Setup

1. Clone the repository or download the script.
2. Run the following command to install the required packages:
    ```bash
    go get github.com/dghubble/oauth1 github.com/joho/godotenv
    ```

3. Create a `.env` file in the root directory with the following variables:

    ```plaintext
    CONSUMER_KEY=your_consumer_key
    CONSUMER_SECRET=your_consumer_secret
    ACCESS_TOKEN=your_access_token
    ACCESS_TOKEN_SECRET=your_access_token_secret
    ```

4. Replace `your_consumer_key`, `your_consumer_secret`, `your_access_token`, `your_access_token_secret`, and `your_bearer_token` with the appropriate values from your Twitter Developer account.

## Running the Application

1. Load the environment variables:
    ```bash
    source .env
    ```

2. Run the Go application:
    ```bash
    go run main.go
    ```

This will post a tweet, wait for 30 seconds, and then delete the tweet.

## Notes

- Ensure your API credentials have the necessary permissions (e.g., `write` and `read` permissions).
- The tweet will be deleted automatically after 30 seconds.

## Dependencies

- `github.com/dghubble/oauth1`: For OAuth1 authentication.
- `github.com/joho/godotenv`: For loading environment variables from the `.env` file.
