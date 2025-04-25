package store

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

type StoreService struct {
	redisClient *redis.Client
}

var (
	storeService = &StoreService{}
	ctx          = context.Background()
)

const CacheDuration = 6 * time.Hour

// StoreInit initializes the Redis client if it hasn't been initialized yet.
func StoreInit() *StoreService {
	if storeService.redisClient != nil {
		// Redis client already initialized, no need to re-initialize
		return storeService
	}

	// Initialize Redis client
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "host.docker.internal:6379", // Redis server address
		Password: "",                          // No password
		DB:       0,                           // Default DB
	})

	// Ping Redis server to ensure the connection works
	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		// Handle Redis connection failure
		fmt.Printf("Failed to connect to Redis: %v\n", err)
		return nil
	}
	fmt.Printf("Redis connection established: %s\n", pong)
	storeService.redisClient = redisClient
	return storeService
}

// SaveUrlMapping saves the short URL and its original URL in Redis.
func SaveUrlMapping(shortUrl string, originalUrl string, userId string) error {
	err := storeService.redisClient.Set(ctx, shortUrl, originalUrl, CacheDuration).Err()
	if err != nil {
		// Return error instead of panic for better error handling
		return fmt.Errorf("failed to save URL mapping for %s: %v", shortUrl, err)
	}
	return nil
}

// RetrieveInitialUrl retrieves the original URL from Redis using the short URL.
func RetrieveInitialUrl(shortUrl string) (string, error) {
	result, err := storeService.redisClient.Get(ctx, shortUrl).Result()
	if errors.Is(err, redis.Nil) {
		// Handle key not found gracefully
		fmt.Printf("Short URL %s not found in Redis\n", shortUrl)
		return "", nil // Return empty string or an appropriate value to indicate not found
	}
	if err != nil {
		// Return error instead of panic
		return "", fmt.Errorf("failed to retrieve initial URL for %s: %v", shortUrl, err)
	}
	return result, nil
}
