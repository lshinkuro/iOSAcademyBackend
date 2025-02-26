package config

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func ConnectRedis() {
	// Get Redis URL from environment variable
	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		log.Println("Warning: REDIS_URL environment variable is not set. Caching will be disabled.")
		return
	}

	// Parse Redis options from URL
	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		log.Printf("Warning: Failed to parse Redis URL: %v. Caching will be disabled.", err)
		return
	}

	// Create Redis client
	RedisClient = redis.NewClient(opt)

	// Test connection with retry
	maxRetries := 3
	for i := 0; i < maxRetries; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		err := RedisClient.Ping(ctx).Err()
		cancel()

		if err == nil {
			log.Println("Successfully connected to Redis")
			return
		}

		if i < maxRetries-1 {
			log.Printf("Failed to connect to Redis (attempt %d/%d): %v. Retrying...", i+1, maxRetries, err)
			time.Sleep(2 * time.Second)
		} else {
			log.Printf("Warning: Failed to connect to Redis after %d attempts: %v. Caching will be disabled.", maxRetries, err)
			CloseRedis()
			RedisClient = nil
		}
	}
}

// Close closes the Redis connection
func CloseRedis() {
	if RedisClient != nil {
		if err := RedisClient.Close(); err != nil {
			log.Printf("Error closing Redis connection: %v", err)
		}
		RedisClient = nil
	}
}
