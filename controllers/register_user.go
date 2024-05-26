package controllers

import (
	"context"
	"nice/helpers"
	"nice/models"

	"github.com/gofiber/fiber/v3"
)

func RegisterUser(c fiber.Ctx) error {
	client := helpers.ConnectToDB()
	defer client.Disconnect(context.TODO())

	user := new(models.Users)

	err := c.Bind().JSON(user)

	helpers.HandleError(err)

	response := models.Repsonse{
		Message: "User Registerd",
		Success: true,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
