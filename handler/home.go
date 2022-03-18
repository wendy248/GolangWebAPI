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
	id := c.Param("halo") //query Param

	c.JSON(http.StatusOK, gin.H{
		"messages": "keterangan barang",
		"id":       id,
	})
}

func ProductInformationByName(c *gin.Context) {
	name := c.Param("nama")

	c.JSON(http.StatusOK, gin.H{
		"messages":     "keterangan barang",
		"product name": name,
	})
}
