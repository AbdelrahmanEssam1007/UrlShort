package main

import (
	"fmt"
	"github.com/AbdelrahmanEssam1007/UrlShort/handler"
	"github.com/AbdelrahmanEssam1007/UrlShort/store"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Route to check the status of the API and show all shortened URLs
	r.GET("/", func(c *gin.Context) {
		shortUrls, err := store.GetAllShortUrls()
		if err != nil {
			c.JSON(500, gin.H{
				"message": "Error retrieving short URLs",
			})
			return
		}

		c.JSON(200, gin.H{
			"message":    "Url Shortener API!!",
			"short_urls": shortUrls,
		})
	})

	// Route to create a short URL
	r.POST("/create-short-url", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	// Route to handle redirection
	r.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortRedirect(c)
	})

	// Initialize the store (Redis connection)
	store.StoreInit()

	// Run the server
	err := r.Run("0.0.0.0:9808")
	if err != nil {
		panic(fmt.Sprintf("Failed to start server: %v", err))
	}
}
