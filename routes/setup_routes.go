package routes

import (
	"nice/controllers"

	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", controllers.GetBlogs)
	app.Post("/write", controllers.WriteBlog)
	app.Get("/blog/:id", controllers.GetBlog)
	app.Post("/register", controllers.RegisterUser)
}
