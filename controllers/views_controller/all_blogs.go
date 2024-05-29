package viewscontroller

import (
	"nice/helpers"
	"nice/views"

	"github.com/gofiber/fiber/v3"
)

func AllBlogs(c fiber.Ctx) error {
	return helpers.Render(c, views.AllBlogs())
}
