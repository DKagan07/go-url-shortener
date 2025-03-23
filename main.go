package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	handler "go-url-shortener/handler"
	"go-url-shortener/store"
)

const HostAndPort = "localhost:9898"

func main() {
	r := gin.Default()

	s := store.InitializeStore()
	handlers := handler.Handlers{
		Store: s,
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Go Url Shortener!",
		})
	})

	r.POST("/shortUrl", func(c *gin.Context) {
		handlers.CreateShortUrl(c)
	})

	r.GET("/:shortUrl", func(c *gin.Context) {
		handlers.HandleShortUrlRedirect(c)
	})

	err := r.Run(HostAndPort)
	if err != nil {
		panic("failed to run")
	}
}
