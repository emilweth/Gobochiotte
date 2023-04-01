package bot

import (
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

var (
	dg *discordgo.Session
)

func StartBot(Token string) {
	var err error
	dg, err = discordgo.New("Bot " + Token)
	if err != nil {
		log.WithError(err).Error("Error creating discord session")
		return
	}
	dg.AddHandler(handleMessage)
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		log.WithError(err).Error("Error while opening connection")
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	log.Info("Discord connected")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	err = dg.Close()
	if err != nil {
		return
	}
}
