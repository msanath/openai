package chat

import (
	"encoding/json"
	"msanath/openai/internal/client"
)

const (
	openAIChatCompletionURL = "https://api.openai.com/v1/chat/completions"
)

type Client struct {
	client client.Client
}

func NewClient(apiKey string) Client {
	return Client{
		client: client.NewClient(openAIChatCompletionURL, apiKey),
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
