package store

import (
	"context"
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

func StoreInit() *StoreService {
	if storeService.redisClient != nil {
		return storeService
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "host.docker.internal:6379", // Redis server address
		Password: "",                          // No password
		DB:       0,                           // Default DB
	})

	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		fmt.Printf("Failed to connect to Redis: %v\n", err)
		return nil
	}
	fmt.Printf("Redis connection established: %s\n", pong)
	storeService.redisClient = redisClient
	return storeService
}

func SaveUrlMapping(shortUrl string, originalUrl string, userId string) error {
	shortUrlKey := "short_url:" + shortUrl
	err := storeService.redisClient.Set(ctx, shortUrlKey, originalUrl, CacheDuration).Err()
	if err != nil {
		return fmt.Errorf("failed to save URL mapping for %s: %v", shortUrl, err)
	}
	return nil
}

func RetrieveInitialUrl(shortUrl string) (string, error) {
	shortUrlKey := "short_url:" + shortUrl
	result, err := storeService.redisClient.Get(ctx, shortUrlKey).Result()
	if err != nil && err != redis.Nil {
		return "", fmt.Errorf("failed to retrieve initial URL for %s: %v", shortUrl, err)
	}
	return result, nil
}

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
