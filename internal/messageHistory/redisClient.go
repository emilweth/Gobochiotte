package messageHistory

import (
	"context"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
	"os"
)

var (
	client *redis.Client
	ctx    = context.Background()
)

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.WithError(err).Panic("Redis connection failed")
	}
	log.Info("Redis connected")
}
