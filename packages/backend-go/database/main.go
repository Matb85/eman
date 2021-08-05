package database

import (
	"os"

	"go.mongodb.org/mongo-driver/bson/primitive"
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

type User struct {
	Id          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email       string             `json:"email"`
	Password    string             `json:"password"`
	FirstName   string             `json:"firstname"`
	LastName    string             `json:"lastname"`
	Enterprises []string           `json:"enterprises"`
	Folder      string             `json:"folder"`
}
