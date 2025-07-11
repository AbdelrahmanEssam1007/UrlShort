package main

import (
	"fmt"
	"os"

	"github.com/AbdelrahmanEssam1007/UrlShort/handler"
	"github.com/AbdelrahmanEssam1007/UrlShort/store"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	// Swagger packages
	_ "github.com/AbdelrahmanEssam1007/UrlShort/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           URL Shortener API
// @version         1.0
// @description     A simple URL shortener built with Go, Gin, and Redis.
// @host            localhost:9808
// @BasePath        /
func main() {
	// Initialize services
	store.StoreInit()

	// Setup and run server
	port := getPort()
	r := setupRouter()
	startServer(r, port)
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	// CORS configuration
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	}))

	// Swagger docs route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API routes
	r.GET("/", listShortenedUrls)
	r.POST("/create-short-url", handler.CreateShortUrl)
	r.GET("/:shortUrl", handler.HandleShortRedirect)

	return r
}

func listShortenedUrls(c *gin.Context) {
	urlMap, err := store.GetAllShortUrls()
	if err != nil {
		c.JSON(500, gin.H{"message": "Error retrieving short URLs"})
		return
	}

	baseURL := os.Getenv("APP_BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:9808"
	}

	var shortUrls []gin.H
	for shortCode, originalUrl := range urlMap {
		shortUrls = append(shortUrls, gin.H{
			"short_url":    fmt.Sprintf("%s/%d", baseURL, shortCode),
			"original_url": originalUrl,
		})
	}

	c.JSON(200, gin.H{
		"message":    "URL Shortener API!!",
		"short_urls": shortUrls,
	})
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9808"
	}
	return port
}

func startServer(r *gin.Engine, port string) {
	address := "0.0.0.0:" + port
	if err := r.Run(address); err != nil {
		panic(fmt.Sprintf("❌ Failed to start server: %v", err))
	}
}
