package main

import (
	"Gobochiotte/internal/bot"
	"flag"
	"os"

	log "github.com/sirupsen/logrus"
)

var (
	DiscordToken string
	verbosity    *string
)

func init() {
	DiscordToken = os.Getenv("DISCORD_TOKEN")
	verbosity = flag.String("verbosity", "info", "Set the verbosity level (trace, debug, info, warn, error, fatal, panic)")
	flag.Parse()
	level, err := log.ParseLevel(*verbosity)
	if err != nil {
		log.Fatalf("Invalid verbosity level: %v", err)
	}

	// Set the Logrus verbosity level
	log.SetLevel(level)
}

func main() {
	bot.StartBot(DiscordToken)
}
