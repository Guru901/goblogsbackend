package routes

import (
	"nice/controllers"
	viewscontroller "nice/controllers/views_controller"

	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", viewscontroller.Home)
	app.Get("/allblogs", viewscontroller.AllBlogs)
	app.Post("/write", controllers.WriteBlog)
	app.Get("/blog/:id", controllers.GetBlog)
	app.Get("/blog", controllers.GetBlogs)
	app.Post("/register", controllers.RegisterUser)
	app.Post("/login", controllers.LoginUser)
	app.Post("/comment", controllers.Comment)
}
