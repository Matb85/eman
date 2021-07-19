package database

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var UserCol *mongo.Collection

func Connect() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:"+os.Getenv("DB_PORT")))
	if err != nil {
		panic(err)
	}

	db := client.Database("test")
	UserCol = db.Collection("users")

}
