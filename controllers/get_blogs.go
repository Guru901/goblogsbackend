package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"nice/helpers"
	"nice/models"

	"github.com/gofiber/fiber/v3"
	"go.mongodb.org/mongo-driver/bson"
)

func GetBlogs(c fiber.Ctx) error {
	client := helpers.ConnectToDB()

	defer client.Disconnect(context.TODO())

	coll := client.Database("goblogs").Collection("blogs")

	cursor, err := coll.Find(context.Background(), bson.M{})
	helpers.HandleError(err)
	var results []models.Blogs
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	// Prints the results of the find operation as structs
	for _, result := range results {
		cursor.Decode(&result)
		output, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", output)
	}

	// reponse := models.Repsonse{
	// 	Success: true,
	// 	Message: "got the blogs",
	// }

	return c.Status(fiber.StatusOK).JSON(results)
}
