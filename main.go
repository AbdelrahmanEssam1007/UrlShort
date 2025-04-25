package main

import (
	"fmt"
	"github.com/AbdelrahmanEssam1007/UrlShort/handler"
	"github.com/AbdelrahmanEssam1007/UrlShort/store"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Url Shortener API!!",
		})
	})

	r.POST("/create-short-url", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})
	r.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortRedirect(c)
	})

	store.StoreInit()

	err := r.Run("0.0.0.0:9808")
	if err != nil {
		panic(fmt.Sprintf("Failed to start server: %v", err))
	}
}
