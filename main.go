package main

import (
	"log"
	"sample/db"
	"sample/models"
	"sample/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

// @Title InstaPay
// @version 1.16.2
//@Basepath

func main() {
	// .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize Fiber app
	app := fiber.New()

	// Database connection
	err = db.ConnectDB()
	if err != nil {
		log.Fatalf("Error connecting to the database %v", err)
	}

	// Auto-migrate models
	db.DB.AutoMigrate(&models.Notifications_Data{})

	// Setup routes
	routes.SetupRoutes(app)
	routes.AuthenticatedRoutes(app)

	// Start server
	err = app.Listen(":1432")
	if err != nil {
		log.Fatal("Error starting the server")
	}
}
