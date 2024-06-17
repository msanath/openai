package main

import (
	"bufio"
	"fmt"
	"msanath/openai/extensions/chatbot"
	"msanath/openai/pkg/chat"
	"msanath/openai/pkg/models"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

const (
	columnLimit = 150
)

func main() {
	apiKey, err := getAPIKey()
	if err != nil {
		fmt.Println(err)
		return
	}

	cb := chatbot.NewClient(chat.NewClient(apiKey), chatbot.WithModel(models.GPT_4))
	chatLoop(cb)
}

func chatLoop(cb chatbot.Client) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("You: ")
		shouldExit := false
		if !scanner.Scan() {
			break
		}
		text := scanner.Text()
		text = strings.TrimSpace(text)
		if text == "goodbye" || text == "exit" {
			shouldExit = true
		}

		// Call the chat function to get the response
		response, err := cb.Send(text)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		response = formatText(response)
		fmt.Println("Bot:", response)
		fmt.Println()

		if shouldExit {
			break
		}
	}
}

func formatText(text string) string {
	var lines []string
	for i := 0; i < len(text); i += columnLimit {
		end := i + columnLimit
		if end > len(text) {
			end = len(text)
		}
		line := text[i:end]
		if end < len(text) && text[end-1] != '\n' && !strings.Contains(line, "\n") {
			lastSpaceIndex := strings.LastIndex(line, " ")
			if lastSpaceIndex != -1 {
				line = line[:lastSpaceIndex] + "\n" + line[lastSpaceIndex+1:]
			} else {
				line = strings.TrimRight(line, " ") + "\n"
			}
		}
		lines = append(lines, line)
	}
	return strings.Join(lines, "")
}

func getAPIKey() (string, error) {
	_ = godotenv.Load()

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("OPENAI_API_KEY is not set")
	}
	return apiKey, nil
}
