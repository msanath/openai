package main

import (
	"fmt"
	"msanath/openai/pkg/chat"
	"msanath/openai/pkg/models"
	"msanath/openai/pkg/roles"
	"os"

	"github.com/joho/godotenv"
)

const (
	openAIURL = "https://api.openai.com/v1/chat/completions"
)

func main() {
	apiKey, err := getAPIKey()
	if err != nil {
		fmt.Println(err)
		return
	}

	chatClient := chat.NewClient(openAIURL, apiKey)

	request := chat.ChatRequest{
		Model: models.GPT_3_5_Turbo,
		Messages: []chat.Message{
			{
				Content: "What is the capital of the United States?",
				Role:    roles.User,
			},
		},
	}

	response, err := chatClient.Send(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(response.Choices[0].Message.Content)
}

func getAPIKey() (string, error) {
	_ = godotenv.Load()

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("OPENAI_API_KEY is not set")
	}
	return apiKey, nil
}
