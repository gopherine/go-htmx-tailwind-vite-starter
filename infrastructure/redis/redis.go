package database

import (
	"context"
	"os"
	"sync"

	"github.com/go-redis/redis/v8"
	"github.com/rs/zerolog/log"
)

var (
	redisClientInstance *redis.Client
	redisOnce           sync.Once
)

// NewRedisClient provides a global point of access to the Redis client.
func NewRedisClient() *redis.Client {
	redisOnce.Do(func() {
		redisClientInstance = redis.NewClient(&redis.Options{
			Addr:     os.Getenv("REDIS_ADDR"),     // replace with your Redis server address
			Password: os.Getenv("REDIS_PASSWORD"), // no password set
			DB:       0,                           // use default DB
		})
		pong, err := redisClientInstance.Ping(context.Background()).Result()
		if err != nil {
			panic(err)
		}

		log.Info().Msgf("Redis connected: %s", pong)
	})

	return redisClientInstance
}

// CloseRedisClient closes the Redis client connection.
func CloseRedisClient() {
	if redisClientInstance != nil {
		err := redisClientInstance.Close()
		if err != nil {
			log.Error().Err(err).Msg("failed to close redis client")
		}
	}
}
