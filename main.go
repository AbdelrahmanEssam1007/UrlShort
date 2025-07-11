package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/AbdelrahmanEssam1007/UrlShort/handler"
	"github.com/AbdelrahmanEssam1007/UrlShort/store"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load environment variables from .env (only in dev/local)
	os.Getenv("REDIS_ADDR")

	// Set up Gin with default middleware
	r := gin.Default()

	// Enable CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // or your frontend domain
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	}))

	// GET: Return all shortened URLs
	r.GET("/", func(c *gin.Context) {
		urlMap, err := store.GetAllShortUrls()
		if err != nil {
			c.JSON(500, gin.H{"message": "Error retrieving short URLs"})
			return
		}

		var shortUrls []gin.H
		baseURL := os.Getenv("APP_BASE_URL")
		if baseURL == "" {
			baseURL = "http://localhost:9808"
		}

		for shortCode, originalUrl := range urlMap {
			shortUrls = append(shortUrls, gin.H{
				"short_url":    baseURL + "/" + strconv.Itoa(shortCode),
				"original_url": originalUrl,
			})
		}

		c.JSON(200, gin.H{
			"message":    "URL Shortener API!!",
			"short_urls": shortUrls,
		})
	})

	// POST: Create a new short URL
	r.POST("/create-short-url", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	// GET: Redirect from short URL to original URL
	r.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortRedirect(c)
	})

	// Initialize Redis
	store.StoreInit()

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "9808"
	}
	if err := r.Run("0.0.0.0:" + port); err != nil {
		panic(fmt.Sprintf("❌ Failed to start server: %v", err))
	}
}
