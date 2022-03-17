package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func QueryFull(c *gin.Context) {
	title := c.Query("judul")
	price := c.Query("harga")

	c.JSON(http.StatusOK, gin.H{
		"messages": "ini function QueryFull dari file item.go",
		"title":    title,
		"price":    price,
	})
}
