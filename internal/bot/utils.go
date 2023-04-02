package bot

import (
	"regexp"

	log "github.com/sirupsen/logrus"
)

var (
	mentionRegex = regexp.MustCompile("<@([0-9]+)>")
)

func GetUsernameForId(discordUserId string) (string, error) {
	user, err := dg.User(discordUserId)
	if err != nil {
		log.WithField("discorduserId", discordUserId).WithError(err).Error("Cannot fetch discord username for given ID")
		return "", err
	}

	return user.Username, nil
}

func ReplaceMentionsByUsername(message string) string {
	return mentionRegex.ReplaceAllStringFunc(message, func(mention string) string {
		id := mentionRegex.ReplaceAllString(mention, "$1")
		username, err := GetUsernameForId(id)
		if err != nil {
			return mention
		}
		return username
	})
}
