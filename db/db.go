package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Client ...
var Client *mongo.Client

// ConnectDatabase ...
func ConnectDatabase() {
	fmt.Println("in ConnectDatabase")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	Client, _ = mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))

	if err := Client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	fmt.Println("mongoDB connected")
}
