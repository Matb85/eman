package database

import (
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"redinnlabs.com/redinn-core/utils"
)

var UserCol *mongo.Collection

var EnterpriseCol *mongo.Collection

func Connect() {
	ctx, cancel := utils.CreateDBContext()
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:"+os.Getenv("DB_PORT")))
	if err != nil {
		panic(err)
	}

	db := client.Database("test")

	UserCol = db.Collection("users")
	EnterpriseCol = db.Collection("enterprises")
}
