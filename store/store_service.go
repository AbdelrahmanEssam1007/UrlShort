package store

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

type StoreService struct {
	redisClient *redis.Client
}

var (
	storeService = &StoreService{}
	ctx          = context.Background()
)

const CacheDuration = 6 * time.Hour

// StoreInit initializes the Redis connection
func StoreInit() *StoreService {
	if storeService.redisClient != nil {
		return storeService
	}

	// Get Redis address from environment variable
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "localhost:6379" // fallback for local dev
	}

	redisPassword := os.Getenv("REDIS_PASSWORD") // optional password
	redisDB := 0

	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       redisDB,
	})

	// Test connection
	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		fmt.Printf("❌ Failed to connect to Redis: %v\n", err)
		return nil
	}
	fmt.Printf("✅ Redis connection established: %s\n", pong)

	storeService.redisClient = redisClient
	return storeService
}

// SaveUrlMapping stores the short URL to original URL mapping
func SaveUrlMapping(shortUrl string, originalUrl string, userId string) error {
	shortUrlKey := "short_url:" + shortUrl
	err := storeService.redisClient.Set(ctx, shortUrlKey, originalUrl, CacheDuration).Err()
	if err != nil {
		return fmt.Errorf("failed to save URL mapping for %s: %v", shortUrl, err)
	}
	return nil
}

// RetrieveInitialUrl retrieves the original URL for a given short URL
func RetrieveInitialUrl(shortUrl string) (string, error) {
	shortUrlKey := "short_url:" + shortUrl
	result, err := storeService.redisClient.Get(ctx, shortUrlKey).Result()
	if err != nil && err != redis.Nil {
		return "", fmt.Errorf("failed to retrieve original URL for %s: %v", shortUrl, err)
	}
	return result, nil
}

// GetAllShortUrls fetches all shortened URLs from Redis
func GetAllShortUrls() ([]string, error) {
	keys, err := storeService.redisClient.Keys(ctx, "short_url:*").Result()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch short URLs: %v", err)
	}

	shortUrls := make([]string, len(keys))
	for i, key := range keys {
		shortUrls[i] = key[len("short_url:"):]
	}

	return shortUrls, nil
}
