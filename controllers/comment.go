package controllers

import (
	"context"
	"fmt"
	"nice/helpers"
	"nice/models"

	"github.com/gofiber/fiber/v3"
)

func Comment(c fiber.Ctx) error {
	client := helpers.ConnectToDB()
	defer client.Disconnect(context.TODO())

	coll := client.Database("goblogs").Collection("comment")

	comment := new(models.Comment)

	err := c.Bind().JSON(comment)

	helpers.HandleError(err)
	result, err := coll.InsertOne(context.TODO(), comment)
	helpers.HandleError(err)
	fmt.Printf("Document inserted with ID: %s\n", result.InsertedID)

	response := models.Repsonse{
		Message: "Comment Added",
		Success: true,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
