package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func PriceCheck(c *gin.Context) {
	title := c.Query("judul")
	price := c.Query("harga")

	c.JSON(http.StatusOK, gin.H{
		"Messages":  "item buku",
		"Item name": title,
		"Time":      time.Now(),
		"Price":     price,
	})
}

func PriceCheckFull(c *gin.Context) {
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

func ProductCondition(c *gin.Context) {
	name := c.Param("name")
	cond := c.Param("kondisi")
	price := c.Query("harga")
	stock := c.Query("jumlah")

	c.JSON(http.StatusOK, gin.H{
		"Messages":     "Ini adalah function ProductCondition",
		"Product name": name,
		"Condition":    cond,
		"Price":        price,
		"Stock":        stock,
	})
}

func ProductQuantity(c *gin.Context) {
	title := c.Query("nama")
	quanty := c.Query("jumlah")
	pos := c.Query("tempat")

	message := "Stok dari " + title + " berjumlah " + quanty + ". Semuanya ada di " + pos

	c.String(http.StatusOK, message)
	// c.JSON(http.StatusOK, gin.H{
	// 	"Messages":  "Ini adalah jumlah stok barang",
	// 	"Item name": title,
	// 	"Quantity":  quanty,
	// })
}
