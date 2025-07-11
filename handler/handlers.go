package handler

import (
	"fmt"
	"net/http"

	"github.com/AbdelrahmanEssam1007/UrlShort/shortener"
	"github.com/AbdelrahmanEssam1007/UrlShort/store"
	"github.com/gin-gonic/gin"
)

type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
	UserId  string `json:"user_id" binding:"required"`
}

// CreateShortUrl handles the creation of short URLs
// @Summary      Create a new short URL
// @Description  Accepts a long URL and user ID and returns a shortened version
// @Tags         Shortener
// @Accept       json
// @Produce      json
// @Param        request body UrlCreationRequest true "Original URL and User ID"
// @Success      200  {object} map[string]string
// @Failure      400  {object} map[string]string
// @Failure      500  {object} map[string]string
// @Router       /create-short-url [post]
func CreateShortUrl(c *gin.Context) {
	var creationRequest UrlCreationRequest
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		fmt.Println("Failed to bind JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("Creating short URL for: %s by %s\n", creationRequest.LongUrl, creationRequest.UserId)
	shortUrl := shortener.GenerateShortUrl(creationRequest.LongUrl, creationRequest.UserId)

	err := store.SaveUrlMapping(shortUrl, creationRequest.LongUrl, creationRequest.UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save URL mapping"})
		return
	}

	host := "http://localhost:9808/"
	c.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": host + shortUrl,
	})
}

// HandleShortRedirect handles redirection from short to original URL
// @Summary      Redirect to original URL
// @Description  Redirects the user from a short URL to the original long URL
// @Tags         Shortener
// @Produce      plain
// @Param        shortUrl path string true "Short URL code"
// @Success      302
// @Failure      404  {object} map[string]string
// @Failure      500  {object} map[string]string
// @Router       /{shortUrl} [get]
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
