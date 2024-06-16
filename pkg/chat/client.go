package chat

import (
	"encoding/json"
	"msanath/openai/internal/client"
)

type Client struct {
	client client.Client
}

func NewClient(url, apiKey string) *Client {
	return &Client{
		client: client.NewClient(url, apiKey),
	}
}

func (c *Client) Send(request ChatRequest) (*ChatResponse, error) {
	var response ChatResponse
	resp, err := c.client.Send(request)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
