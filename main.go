package main

import (
	"inventory-ms/config"
	"inventory-ms/db"
	"inventory-ms/routes"
	"log"
)

func main() {
	// Load configuration
	config.LoadConfig()

	// Initialize database connection
	db.InitDB()

	// Initialize router
	router := routes.SetupRouter()

	// Start the server on the configured port
	port := config.Config.Server.Port
	log.Printf("Inventory Service is running on port %s...", port)
	router.Run(":" + port)
}
