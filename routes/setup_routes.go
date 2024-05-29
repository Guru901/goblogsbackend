package routes

import (
	"nice/controllers"
	viewscontroller "nice/controllers/views_controller"

	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", viewscontroller.Home)
	app.Post("/write", controllers.WriteBlog)
	app.Get("/blog/:id", controllers.GetBlog)
	app.Post("/register", controllers.RegisterUser)
	app.Post("/login", controllers.LoginUser)
}
