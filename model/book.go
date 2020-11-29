package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Book struct
type Book struct {
	ID     primitive.ObjectID `json:"id" bson:"id"`
	Isbn   string             `json:"isbn" bson:"isbn"`
	Title  string             `json:"title" bson:"title"`
	Author *Author            `json:"author" bson:"author"`
}
