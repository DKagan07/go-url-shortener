package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"go-url-shortener/shortener"
	"go-url-shortener/store"
)

type Handlers struct {
	Store *store.StorageService
}

type CreationRequest struct {
	LongUrl string `json:"longUrl" binding:"required"`
	UserId  string `json:"userId"  binding:"required"`
}

func (h *Handlers) HandleShortUrlRedirect(c *gin.Context) {}

func (h *Handlers) CreateShortUrl(c *gin.Context) {
	var creationRequest CreationRequest
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	fmt.Printf("%+v\n", creationRequest)

	shortUrl := shortener.GenerateShortLink(creationRequest.LongUrl, creationRequest.UserId)
	h.Store.SaveUrlMapping(shortUrl, creationRequest.LongUrl, creationRequest.UserId)
}
