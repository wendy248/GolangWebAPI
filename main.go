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

	//localhost:8080/item?judul=filosofi-teras&harga=20000 (soal 1)
	r.GET("/item", handler.PriceCheck)

	//localhost:8080/id/2/product?name="tv"&price=1000 (soal 2)
	r.GET("/id/:number/product", handler.PriceCheckFull)

	//localhost:8080/customer/wendy/10/surabaya
	r.GET("/customer/:nama/:id/:kota", handler.UserInformation)

	//localhost:8080/stock?nama=baju&jumlah=100&tempat=surabaya
	r.GET("/stock", handler.ProductQuantity)

	//localhost:8080/QuantyStock/baju/bekas?harga=1000&jumlah=100
	r.GET("/QuantyStock/:name/:kondisi", handler.ProductCondition)

	// r.POST("/mahasiswa", handler.MahasiswaHandler)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
