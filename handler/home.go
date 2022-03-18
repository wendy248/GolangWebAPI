package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "this is file home.go",
		"status":  "oke",
	})
}

func ProductInformation(c *gin.Context) {
	name := c.Param("nama")
	id := c.Param("id") //query Param
	price := c.Param("harga")

	c.JSON(http.StatusOK, gin.H{
		"messages":     "Ini function ProductInformation",
		"Product name": name,
		"id":           id,
		"Price":        price,
	})
}
