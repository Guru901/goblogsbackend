package controllers

import (
	"context"
	"fmt"
	"nice/helpers"
	"nice/models"

	"github.com/gofiber/fiber/v3"
)

func WriteBlog(c fiber.Ctx) error {
	client := helpers.ConnectToDB()

	defer client.Disconnect(context.Background())

	coll := client.Database("goblogs").Collection("blogs")

	new_blog := new(models.Blogs)

	err := c.Bind().Body(new_blog)
	helpers.HandleError(err)

	result, err := coll.InsertOne(context.TODO(), new_blog)
	helpers.HandleError(err)
	fmt.Printf("Document inserted with ID: %s\n", result.InsertedID)

	response := models.Repsonse{
		Message: "Blog Written",
		Success: true,
	}
	return c.Status(fiber.StatusOK).JSON(response)
}
