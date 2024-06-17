// This is an example of how to use the chat package to interact with OpenAI's chat API.
// It demonstrates how to use tools in the chat API to call a (mock) function and get the stock value of a company.
package main

import (
	"encoding/json"
	"fmt"
	"github.com/msanath/openai/pkg/chat"
	"github.com/msanath/openai/pkg/models"
	"github.com/msanath/openai/pkg/roles"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	request := chat.ChatRequest{
		Model: models.GPT_4,
		Messages: []chat.Message{
			{
				Content: "What is price of TSLA?",
				Role:    roles.User,
			},
		},
		Tools: []chat.Tool{
			{
				Type: chat.ToolTypeFunction,
				Function: chat.FunctionDefinition{
					Description: "Get the value of stock",
					Name:        "GetStockValue",
					Parameters: json.RawMessage(`{
						"type": "object",
						"properties": {
							"symbol": {
								"type": "string",
								"description": "The stock symbol, e.g. MSFT"
							},
							"competitorSymbol": {
								"type": "string",
								"description": "The symbol of the competitor closest to the symbol stock, e.g. AAPL"
							}
						},
						"required": ["symbol", "competitorSymbol"]
					}`),
				},
			},
		},
	}

	apiKey, err := getAPIKey()
	if err != nil {
		fmt.Println(err)
		return
	}
	chatClient := chat.NewClient(apiKey)
	response, err := chatClient.Send(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, toolCall := range response.Choices[0].Message.ToolCalls {
		if toolCall.Type == chat.ToolTypeFunction {
			if toolCall.FunctionResponse.Name == "GetStockValue" {

				var args map[string]string
				err := json.Unmarshal([]byte(toolCall.FunctionResponse.Arguments), &args)
				if err != nil {
					fmt.Println("Error unmarshalling arguments:", err)
					return
				}

				stock := args["symbol"]
				comp := args["competitorSymbol"]

				GetStockValue(stock, comp)
			} else {
				fmt.Println("Function not found")
				return
			}
		}
	}
}

func GetStockValue(stock, comp string) {
	fmt.Println("The stock value of", stock, "is $100. Its main competitor is", comp)
}

func getAPIKey() (string, error) {
	_ = godotenv.Load()

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("OPENAI_API_KEY is not set")
	}
	return apiKey, nil
}
