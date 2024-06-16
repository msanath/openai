// chatbot is a package that implements a chatbot which stores the historical messages and can respond to new messages.
package chatbot

import (
	"msanath/openai/pkg/chat"
	"msanath/openai/pkg/models"
	"msanath/openai/pkg/roles"
)

type options struct {
	model   models.Model
	persona string
}

type Option func(*options)

func WithModel(model models.Model) Option {
	return func(o *options) {
		o.model = model
	}
}

func WithPersona(persona string) Option {
	return func(o *options) {
		o.persona = persona
	}
}

type Client struct {
	client  chat.Client
	request chat.ChatRequest
	opts    options
}

func NewClient(c chat.Client, opts ...Option) Client {
	o := options{
		model:   models.GPT_3_5_Turbo,
		persona: "You are a helpful assistant.",
	}
	for _, opt := range opts {
		opt(&o)
	}

	// Set up a chat request.
	request := chat.ChatRequest{
		Model: o.model,
		Messages: []chat.Message{
			{
				Content: o.persona,
				Role:    roles.System,
			},
		},
	}

	return Client{
		client:  c,
		opts:    o,
		request: request,
	}
}

func (c *Client) Send(message string) (string, error) {
	// Add the user message to the chat request.
	c.request.Messages = append(c.request.Messages, chat.Message{
		Content: message,
		Role:    roles.User,
	})

	// Send the chat request.
	response, err := c.client.Send(c.request)
	if err != nil {
		return "", err
	}

	// Add the response message to the chat request.
	c.request.Messages = append(c.request.Messages, chat.Message{
		Content: response.Choices[0].Message.Content,
		Role:    roles.System,
	})

	return response.Choices[0].Message.Content, nil
}
