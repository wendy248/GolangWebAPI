package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type BiodataInput struct {
	Nama     string  `json:"nama" binding:"required"`
	Umur     int8    `json:"umur" binding:"required"`
	Domisili string  `json:"domisili" binding:"required"`
	Berat    float32 `json:"berat" binding:"required"`
	Tinggi   float32 `json:"tinggi" binding:"required"`
}

func BiodataHandler(c *gin.Context) {
	var biodataInput BiodataInput

	err := c.ShouldBindJSON(&biodataInput)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error %s, message: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message":      "Berhasil input data",
		"Waktu Input":         time.Now(),
		"Nama":         biodataInput.Nama,
		"Umur":         biodataInput.Umur,
		"Domisili":     biodataInput.Domisili,
		"Berat Badan":  biodataInput.Berat,
		"Tinggi Badan": biodataInput.Tinggi,
	})
}
