package bot

import (
	log "github.com/sirupsen/logrus"
	"regexp"
	"strings"
)

func GetUsernameForId(discordUserId string) (string, error) {
	user, err := dg.User(discordUserId)
	if err != nil {
		log.WithField("discorduserId", discordUserId).WithError(err).Error("Cannot fetch discord username for given ID")
		return "", err
	}

	return user.Username, nil
}

func ReplaceMentionsByUsername(message string) (string, error) {
	mentionRegex := regexp.MustCompile("<@([0-9]+)>")

	replacer := func(mention string) string {
		id := mentionRegex.ReplaceAllString(mention, "$1")
		username, err := GetUsernameForId(id)
		if err != nil {
			return mention
		}
		return username
	}

	mentions := mentionRegex.FindAllString(message, -1)
	for _, mention := range mentions {
		message = strings.Replace(message, mention, replacer(mention), 1)
	}

	return message, nil
}
