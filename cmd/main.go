package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"go-fiber/database"
	"go-fiber/routes"
)

func main() {

	//SETUP FIBER
	app := fiber.New(fiber.Config{
		Views: html.New("./views", ".html"),
	})

	//ACCESS FOLDER PUBLIC TO FIBER FRAMEWORK
	app.Static("/", "public")

	// Initialize MySQL Connection
	database.ConnectMySQL()

	// Setup routes
	routes.SetupRoutes(app)

	// Start the Fiber app
	app.Listen(":3000")

}
