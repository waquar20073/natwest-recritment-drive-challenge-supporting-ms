package routes

import (
	"inventory-ms/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/inventory/:id/availability", handlers.CheckAvailability)
	router.POST("/inventory", handlers.CreateInventoryHandler)
	router.GET("/inventory", handlers.GetAllInventoryHandler)
	router.GET("/inventory/:id", handlers.GetInventoryByIDHandler)
	router.PUT("/inventory/:id", handlers.UpdateInventoryHandler)
	router.DELETE("/inventory/:id", handlers.DeleteInventoryHandler)
	router.POST("/inventory/deduct", handlers.BulkDeductInventoryHandler)
	return router
}
