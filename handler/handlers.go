package handler

import (
	"fmt"
	"github.com/AbdelrahmanEssam1007/UrlShort/shortener"
	"github.com/AbdelrahmanEssam1007/UrlShort/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
	UserId  string `json:"user_id" binding:"required"`
}

func CreateShortUrl(c *gin.Context) {
	var creationRequest UrlCreationRequest
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortUrl := shortener.GenerateShortUrl(creationRequest.LongUrl, creationRequest.UserId)
	store.SaveUrlMapping(shortUrl, creationRequest.LongUrl, creationRequest.UserId)

	host := "http://localhost:9808/"
	c.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": host + shortUrl,
	})

}

func HandleShortRedirect(c *gin.Context) {
	short := c.Param("shortUrl")
	originalUrl, err := store.RetrieveInitialUrl(short)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve original URL"})
		return
	}
	if originalUrl == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Short URL '%s' not found", short)})
		return
	}
	c.Redirect(http.StatusFound, originalUrl)
}
