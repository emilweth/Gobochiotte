package main

import (
	"Gobochiotte/internal/bot"
	"os"
)

var (
	DiscordToken string
	OpenAiToken  string
)

func init() {
	DiscordToken = os.Getenv("DISCORD_TOKEN")
	OpenAiToken = os.Getenv("OPENAI_TOKEN")
}

func main() {
	bot.StartBot(DiscordToken)
}
