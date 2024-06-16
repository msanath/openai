// An implementation of the request and response schema for the chat service
// Reference - https://platform.openai.com/docs/api-reference/chat
package chat

import (
	"msanath/openai/pkg/models"
	"msanath/openai/pkg/roles"
)

// ChatRequest is the request schema for the chat service.
type ChatRequest struct {
	Model    models.Model `json:"model"`       // ID of the model to use.
	Messages []Message    `json:"messages"`    // A list of messages comprising the conversation so far.
	N        int          `json:"n,omitempty"` // How many chat completion choices to generate for each input message.
}

// Message is a list messages comprising the conversation so far.
type Message struct {
	Content string     `json:"content"`        // The contents of the system message.
	Role    roles.Role `json:"role"`           // The role of the messages author
	Name    string     `json:"name,omitempty"` // An optional name for the participant
}

// ChatResponse is the response schema for the chat service.
type ChatResponse struct {
	ID      string       `json:"id"`      // A unique identifier for the chat completion.
	Object  string       `json:"object"`  // The object type, which is always chat.completion.
	Created int          `json:"created"` // The Unix timestamp (in seconds) of when the chat completion was created
	Model   models.Model `json:"model"`   // The model used for the chat completion
	Choices []Choice     `json:"choices"` // A list of chat completion choices. Can be more than one if `n` is greater than 1.
	Usage   Usage        `json:"usage"`   // Usage statistics for the completion request.
}

// Choice represents a choice in a conversation.
type Choice struct {
	Index        int     `json:"index"`         // Index of the choice.
	Message      Message `json:"message"`       // The message associated with the choice.
	FinishReason string  `json:"finish_reason"` // The reason for finishing the conversation.
}

// Usage statistics for the completion request.
type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`     // Number of tokens in the prompt.
	CompletionTokens int `json:"completion_tokens"` // Number of tokens in the generated completion.
	Total            int `json:"total"`             // Total number of tokens used in the request (prompt + completion).
}
