package main

import (
	"log"

	"sample/db"
	"sample/models"
	"sample/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

//@Title InstaPay
//@version 1.16.2

func main() {

	// .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//database connection
	err = db.ConnectDB()
	if err != nil {
		log.Fatalf("Error connecting to the database %v", err)
	}

	db.DB.AutoMigrate(&models.Notifications_Data{})
	//fiber app

	app := fiber.New()

	//connecting routes
	routes.SetupRoutes(app)

	//start server
	err = app.Listen(":8080")
	if err != nil {
		log.Fatal("Error starting the server")
	}
}
