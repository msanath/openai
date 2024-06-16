package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	url    string
	apiKey string
}

// NewClient creates a new client with the given URL and API key.
func NewClient(url, apiKey string) Client {
	return Client{url: url, apiKey: apiKey}
}

func (c *Client) Send(payload interface{}) ([]byte, error) {
	var result []byte

	// Convert the payload to JSON
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return result, fmt.Errorf("error marshalling JSON: %w", err)
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", c.url, bytes.NewBuffer(jsonData))
	if err != nil {
		return result, fmt.Errorf("error creating request: %w", err)
	}

	// Set the appropriate headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return result, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	// Check the response status
	if resp.StatusCode != http.StatusOK {
		return result, fmt.Errorf("received non-200 response status: %s", resp.Status)
	}

	return io.ReadAll(resp.Body)
}
