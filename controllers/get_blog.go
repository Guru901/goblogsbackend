package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"nice/helpers"
	"nice/models"

	"github.com/gofiber/fiber/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetBlog(c fiber.Ctx) error {
	var param string = c.Params("id")
	fmt.Println(param)

	client := helpers.ConnectToDB()
	defer client.Disconnect(context.TODO())

	coll := client.Database("goblogs").Collection("blogs")
	fmt.Println(coll)
	filter := bson.D{{Key: "title", Value: "nice"}}
	var result models.Blogs
	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	// Prints a message if no documents are matched or if any
	// other errors occur during the operation
	if err != nil {
		if err == mongo.ErrNoDocuments {
			panic(err)
		}
		panic(err)
	}
	// end findOne

	output, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", output)

	response := models.Repsonse{
		Message: "got the blog",
		Success: true,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

