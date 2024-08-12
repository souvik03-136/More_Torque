package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/souvik03-136/more-torque/internal/service"
)

func DecodeVehicle(c *gin.Context) {
	vin := c.Param("vin")
	data, err := service.DecodeVIN(vin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": data})
}

func AddVehicle(c *gin.Context) {
	var input struct {
		VIN string `json:"vin"`
		Org string `json:"org"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	data, err := service.AddVehicle(input.VIN, input.Org)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": data})
}

func GetVehicle(c *gin.Context) {
	vin := c.Param("vin")
	data, err := service.GetVehicle(vin)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": data})
}
