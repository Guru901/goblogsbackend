package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Comment struct {
	ByUser  primitive.ObjectID `json:"byUser" bson:"byUser"`
	OnBlog  primitive.ObjectID `json:"onBlog" bson:"onBlog"`
	Content string             `json:"content" bson:"content"`
}

func CommentModel() Comment {
	return Comment{}
}
