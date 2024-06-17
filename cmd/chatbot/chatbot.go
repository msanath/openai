package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/msanath/openai/extension/chatbot"
	"github.com/msanath/openai/pkg/chat"
	"github.com/msanath/openai/pkg/models"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/theckman/yacspin"
)

const (
	columnLimit = 150
)

var allowedModels = []string{
	string(models.GPT_3_5_Turbo),
	string(models.GPT_4),
	string(models.GPT_4_Turbo),
	string(models.GPT_4o),
}

type options struct {
	model models.Model
}

func newRootCmd() *cobra.Command {
	o := &options{}
	var model string

	cmd := &cobra.Command{
		Use:   "chai",
		Short: "chai is a chatbot that provides a chatbot experience using OpenAI's models.",
		RunE: func(cmd *cobra.Command, args []string) error {
			for _, allowedModel := range allowedModels {
				if model == allowedModel {
					o.model = models.Model(model)
				}
			}

			if o.model == "" {
				return fmt.Errorf("invalid model %s. Allowed models are %v", model, allowedModels)
			}
			return o.Run(cmd.Context())
		},
	}

	cmd.Flags().StringVarP(&model, "model", "m", string(models.GPT_3_5_Turbo), "The openai model to use. Allowed values are "+strings.Join(allowedModels, ", "))
	return cmd
}

func (o *options) Run(ctx context.Context) error {
	fmt.Println("Running with model:", o.model)

	apiKey, err := getAPIKey()
	if err != nil {
		return err
	}

	cb := chatbot.NewClient(chat.NewClient(apiKey), chatbot.WithModel(models.GPT_4))
	return chatLoop(cb)
}

func chatLoop(cb chatbot.Client) error {
	scanner := bufio.NewScanner(os.Stdin)

	cfg := yacspin.Config{
		Frequency:       100 * time.Millisecond,
		CharSet:         yacspin.CharSets[59],
		SuffixAutoColon: true,
		Message:         "communicating with the openai server",
		StopCharacter:   "",
	}
	spinner, err := yacspin.New(cfg)
	if err != nil {
		return err
	}

	for {
		fmt.Print("You: ")
		if !scanner.Scan() {
			break
		}
		text := scanner.Text()
		text = strings.TrimSpace(text)
		if text == "goodbye" || text == "exit" {
			fmt.Println("Bot: Goodbye!")
			break
		}
		_ = spinner.Start()

		// Call the chat function to get the response
		response, err := cb.Send(text)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		_ = spinner.Stop()

		response = formatText(response)
		fmt.Println("Bot:", response)
		fmt.Println()
	}
	return nil
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
