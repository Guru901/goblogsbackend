package controllers

import (
	"context"
	"fmt"
	"nice/helpers"
	"nice/models"

	"github.com/gofiber/fiber/v3"
)

func RegisterUser(c fiber.Ctx) error {
	client := helpers.ConnectToDB()
	defer client.Disconnect(context.TODO())

	coll := client.Database("goblogs").Collection("users")

	user := new(models.Users)

	err := c.Bind().JSON(user)

	helpers.HandleError(err)
	result, err := coll.InsertOne(context.TODO(), user)
	helpers.HandleError(err)
	fmt.Printf("Document inserted with ID: %s\n", result.InsertedID)

	response := models.Repsonse{
		Message: "User Registerd",
		Success: true,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
