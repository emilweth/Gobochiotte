package bot

import (
	"Gobochiotte/internal/messageHistory"
	"Gobochiotte/internal/openaiClient"
	"strings"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"

	"fmt"
)

func handleMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	log.WithField("Message", fmt.Sprintf("%s : %s", m.Author.Username, m.Content)).Debug("New message")

	// Save the message
	saveMessage(s, m)

	// Check if bot is mentionned
	mention := fmt.Sprintf("<@%s>", s.State.User.ID)
	if strings.Contains(m.Content, mention) {
		if m.Author.ID == s.State.User.ID {
			return
		}
		err := replyToMessage(m)
		if err != nil {
			log.WithError(err).Panic("Error while sending reply")
		}
	}

}

func saveMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	log.WithField("Message", fmt.Sprintf("%s : %s", m.Author.Username, m.Content)).Debug("Saving message to redis")
	formattedMessage := ReplaceMentionsByUsername(m.Content)
	toSaveMessage := messageHistory.SavedMessage{
		UserID:   m.Author.ID,
		Username: m.Author.Username,
		Text:     formattedMessage,
		IsSelf:   m.Author.ID == s.State.User.ID,
	}
	err := messageHistory.SaveMessage(m.ChannelID, toSaveMessage)
	if err != nil {
		log.WithError(err).WithField("message", toSaveMessage).Panic("Cannot save message")
	}
}

func replyToMessage(m *discordgo.MessageCreate) error {
	messages, err := messageHistory.GetChannelLastMessages(m.ChannelID)
	if err != nil {
		return err
	}
	err = dg.ChannelTyping(m.ChannelID)
	if err != nil {
		log.WithError(err).Error("Error while sending typing status")
		return err
	}
	reply, err := openaiClient.GenerateResponse(messages)
	if err != nil {
		return err
	}

	reply = strings.ReplaceAll(reply, "everyone", "tout le monde")
	reply = strings.ReplaceAll(reply, "here", "ici")

	_, err = dg.ChannelMessageSendReply(m.ChannelID, reply, m.Reference())
	if err != nil {
		return err
	}

	return nil
}
