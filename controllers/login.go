package controllers

import (
	"context"
	"fmt"
	"nice/helpers"
	"nice/models"

	"github.com/gofiber/fiber/v3"
	"go.mongodb.org/mongo-driver/bson"
)

func LoginUser(c fiber.Ctx) error {
	client := helpers.ConnectToDB()
	defer client.Disconnect(context.TODO())

	coll := client.Database("goblogs").Collection("users")

	user := new(models.Users)

	err := c.Bind().JSON(user)
	filter := bson.D{{"username", user.Username}}
	helpers.HandleError(err)
	var result models.Users
	err = coll.FindOne(context.TODO(), filter).Decode(&result)
	helpers.HandleError(err)

	fmt.Println(result)

	response := models.Repsonse{
		Message: "User Registerd",
		Success: true,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
