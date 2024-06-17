// An implementation of the request and response schema for the chat service
// Reference - https://platform.openai.com/docs/api-reference/chat
package chat

import (
	"encoding/json"
	"msanath/openai/pkg/models"
	"msanath/openai/pkg/roles"
)

// ChatRequest is the request schema for the chat service.
type ChatRequest struct {
	Model    models.Model `json:"model"`           // ID of the model to use.
	Messages []Message    `json:"messages"`        // A list of messages comprising the conversation so far.
	N        int          `json:"n,omitempty"`     // How many chat completion choices to generate for each input message.
	Tools    []Tool       `json:"tools,omitempty"` // A list of tools the model may call..
}

// Message is a list messages comprising the conversation so far.
type Message struct {
	Content   string     `json:"content"`              // The contents of the system message.
	Role      roles.Role `json:"role"`                 // The role of the messages author
	Name      string     `json:"name,omitempty"`       // An optional name for the participant
	ToolCalls []ToolCall `json:"tool_calls,omitempty"` // A list of tool calls made by the model in response to the message.
}

type ToolType string // The type of the tool.

type Tool struct {
	Type     ToolType           `json:"type"`     // The type of the tool. Currently, the only supported type is "function".
	Function FunctionDefinition `json:"function"` // The function to call.
}

const (
	ToolTypeFunction ToolType = "function"
)

type ToolCall struct {
	ID               string           `json:"id"`       // The ID of the tool call.
	Type             ToolType         `json:"type"`     // The type of the tool. Currently, the only supported type is "function".
	FunctionResponse FunctionResponse `json:"function"` // The response from the function.
}

type FunctionResponse struct {
	Name      string `json:"name"`      // The name of the function.
	Arguments string `json:"arguments"` // The arguments passed to the function.
}

type FunctionDefinition struct {
	Description string          `json:"description,omitempty"` // A description of what the function does, used by the model to choose when and how to call the function.
	Name        string          `json:"name"`                  // The name of the function. Must be a-z, A-Z, 0-9, or contain underscores and dashes, with a maximum length of 64.
	Parameters  json.RawMessage `json:"parameters,omitempty"`  // The parameters the functions accepts, described as a JSON Schema object. https://json-schema.org/understanding-json-schema
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
