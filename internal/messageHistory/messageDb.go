package messageHistory

import (
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"

	"fmt"
)

type SavedMessage struct {
	UserID   string `json:"UserID"`
	Username string `json:"username"`
	Text     string `json:"text"`
	IsSelf   bool   `json:"is_self"`
}

func SaveMessage(channelId string, message SavedMessage) error {
	key := fmt.Sprintf("channel:%s:message", channelId)
	score := float64(time.Now().Unix())

	messageJSON, err := json.Marshal(message)
	if err != nil {
		return err
	}

	// Add the message to the sorted set
	if err := client.ZAdd(ctx, key, redis.Z{Score: score, Member: string(messageJSON)}).Err(); err != nil {
		return err
	}

	// Keep only the 10 last messages
	if err := client.ZRemRangeByRank(ctx, key, 0, -11).Err(); err != nil {
		return err
	}

	return nil
}

func GetChannelLastMessages(channelId string) ([]SavedMessage, error) {
	key := fmt.Sprintf("channel:%s:message", channelId)

	jsonMessages, err := client.ZRange(ctx, key, 0, 9).Result()
	if err != nil {
		return nil, err
	}

	messages := make([]SavedMessage, len(jsonMessages))
	for i, jsonMessage := range jsonMessages {
		var message SavedMessage
		if err := json.Unmarshal([]byte(jsonMessage), &message); err != nil {
			return nil, err
		}
		messages[i] = message
	}

	return messages, nil
}
