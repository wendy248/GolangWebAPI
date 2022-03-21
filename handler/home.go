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

func UserInformation(c *gin.Context) {
	name := c.Param("nama")
	id := c.Param("id") //query Param
	city := c.Param("kota")

	salam := "Hai " + name + " dengan id : " + id + ". Anda berasal dari " + city + "."
	c.String(http.StatusOK, salam)
}
