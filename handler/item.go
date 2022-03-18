package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func QueryFull(c *gin.Context) {
	title := c.Query("judul")
	price := c.Query("harga")

	c.JSON(http.StatusOK, gin.H{
		"Messages":  "item buku",
		"Item name": title,
		"Time":      time.Now(),
		"Price":     price,
	})
}

func IDProduct(c *gin.Context) {
	number := c.Param("number")
	name := c.Query("name")
	price := c.Query("price")

	c.JSON(http.StatusOK, gin.H{
		"ID":           number,
		"Messages":     "item name",
		"Product name": name,
		"Price":        price,
	})
}

func PriceStock(c *gin.Context) {
	name := c.Param("name")
	price := c.Query("harga")
	stock := c.Query("jumlah")

	c.JSON(http.StatusOK, gin.H{
		"Messages":     "Ini adalah function PriceStock",
		"Product name": name,
		"Price":        price,
		"Stock":        stock,
	})
}

func Quantity(c *gin.Context) {
	title := c.Query("nama")
	quanty := c.Query("jumlah")

	c.JSON(http.StatusOK, gin.H{
		"Messages":  "Ini adalah jumlah stok barang",
		"Item name": title,
		"Quantity":  quanty,
	})
}
