package main

import (
	"Github/GolangWebAPI/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	r.GET("/home", handler.HomeHandler)
	r.GET("/item", handler.QueryFull) //Nomor 1 filosofi kopi
	r.GET("/id/:number/product", handler.IDProduct)

	r.GET("/id/:a", handler.ProductInformation)
	// r.GET("/id/", handler.ProductInformationByName)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
