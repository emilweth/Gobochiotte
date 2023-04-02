package openaiClient

import (
	"Gobochiotte/internal/messageHistory"
	"context"
	"github.com/mrz1836/go-sanitize"
	"github.com/sashabaranov/go-openai"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)

var (
	openAIClient *openai.Client
)

func init() {
	openAIClient = openai.NewClient(os.Getenv("OPENAI_TOKEN"))
}

func GenerateResponse(discordMessageHistory []messageHistory.SavedMessage) (string, error) {

	var messages []openai.ChatCompletionMessage
	messages = append(messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleSystem,
		Content: "Tu es Robochiotte. Un assistant présent sur le Discord de MisterMV qui se fait passer pour un humain. Personne ne doit savoir que tu es un robot.",
		Name:    "Robochiotte",
	})
	for _, message := range discordMessageHistory {
		role := openai.ChatMessageRoleUser
		if message.IsSelf {
			role = openai.ChatMessageRoleAssistant
		}
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    role,
			Content: message.Text,
			Name:    strings.TrimSpace(sanitize.Alpha(message.Username, true)),
		})
	}

	log.Debug("Querying OpenAI")
	resp, err := openAIClient.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo,
			Messages: messages,
		},
	)

	if err != nil {
		log.WithError(err).Error("ChatCompletion error")
		return "", err
	}

	log.WithField("response", resp.Choices[0].Message.Content).Debug("OpenAI generated response")

	return resp.Choices[0].Message.Content, nil
}
