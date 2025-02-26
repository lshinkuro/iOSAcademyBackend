package cache

import (
	"context"
	"encoding/json"
	"time"

	"course-api/config"

	"github.com/redis/go-redis/v9"
)

// DefaultExpiration is the default cache expiration time
const DefaultExpiration = 15 * time.Minute

// Set stores a value in the cache with the given key and expiration
func Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	if config.RedisClient == nil {
		return nil // Silently skip if Redis is not available
	}
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return config.RedisClient.Set(ctx, key, data, expiration).Err()
}

// Get retrieves a value from the cache and unmarshals it into the provided interface
func Get(ctx context.Context, key string, dest interface{}) error {
	if config.RedisClient == nil {
		return redis.Nil // Return cache miss if Redis is not available
	}
	data, err := config.RedisClient.Get(ctx, key).Bytes()
	if err != nil {
		return err
	}
	return json.Unmarshal(data, dest)
}

// Delete removes a key from the cache
func Delete(ctx context.Context, key string) error {
	if config.RedisClient == nil {
		return nil // Silently skip if Redis is not available
	}
	return config.RedisClient.Del(ctx, key).Err()
}

// Clear removes all keys from the cache
func Clear(ctx context.Context) error {
	if config.RedisClient == nil {
		return nil // Silently skip if Redis is not available
	}
	return config.RedisClient.FlushAll(ctx).Err()
}

// GetOrSet retrieves a value from cache or sets it if not found
func GetOrSet(ctx context.Context, key string, dest interface{}, fetchFunc func() (interface{}, error)) error {
	// Try to get from cache first
	err := Get(ctx, key, dest)
	if err == nil {
		return nil
	}

	// If not in cache or Redis is not available, fetch the data
	data, err := fetchFunc()
	if err != nil {
		return err
	}

	// Store in cache (will be skipped if Redis is not available)
	if err := Set(ctx, key, data, DefaultExpiration); err != nil {
		return err
	}

	// Copy the fresh data to destination
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return json.Unmarshal(jsonData, dest)
}
