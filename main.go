package main

import (
	"log"
	"nice/helpers"
	"nice/routes"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("./.env.local")
	helpers.HandleError(err)

	app := fiber.New()

	app.Use(cors.New())

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
