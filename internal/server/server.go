package server

import (
	"github.com/gin-gonic/gin"
	"github.com/souvik03-136/more-torque/internal/handler"
)

func NewServer() *gin.Engine {
	router := gin.Default()

	// Define your routes and handlers here
	router.GET("/vehicles/decode/:vin", handler.DecodeVehicle)
	router.POST("/vehicles", handler.AddVehicle)
	router.GET("/vehicles/:vin", handler.GetVehicle)
	router.POST("/orgs", handler.CreateOrg)
	router.PATCH("/orgs", handler.UpdateOrg)
	router.GET("/orgs", handler.ListOrgs)

	return router
}
