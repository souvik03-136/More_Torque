package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/souvik03-136/more-torque/internal/service"
)

func ListOrgs(c *gin.Context) {
	orgs, err := service.GetOrgs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": orgs})
}

// Existing handlers should be here
func CreateOrg(c *gin.Context) {
	var input struct {
		Name                    string `json:"name"`
		Account                 string `json:"account"`
		Website                 string `json:"website"`
		FuelReimbursementPolicy string `json:"fuelReimbursementPolicy"`
		SpeedLimitPolicy        string `json:"speedLimitPolicy"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if input.FuelReimbursementPolicy == "" {
		input.FuelReimbursementPolicy = "1000"
	}

	data, err := service.CreateOrg(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": data})
}

func UpdateOrg(c *gin.Context) {
	var input struct {
		Account                 string `json:"account"`
		Website                 string `json:"website"`
		FuelReimbursementPolicy string `json:"fuelReimbursementPolicy"`
		SpeedLimitPolicy        string `json:"speedLimitPolicy"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	data, err := service.UpdateOrg(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": data})
}
