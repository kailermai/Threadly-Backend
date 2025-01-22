package main

import (
	"log"

	"main/database"
	"main/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()
	app := fiber.New()

	app.Use((cors.New(cors.Config{
		AllowCredentials: true, // allows authetication of cookies
		AllowOrigins:     "http://localhost:5173",
		AllowMethods:     "GET, POST, PUT, DELETE",
	})))

	routes.Setup(app)
	log.Fatal(app.Listen(":8000"))
}
