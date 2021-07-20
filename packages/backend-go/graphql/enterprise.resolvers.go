package graphql

import (
	"context"
	"errors"

	"github.com/graph-gophers/graphql-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"redinnlabs.com/redinn-core/auth"
	"redinnlabs.com/redinn-core/database"
)

type Employee struct {
	Ref         graphql.ID `json:"ref"`
	Permissions string     `json:"permissions"`
}
type Address struct {
	Country string `json:"country"`
	Zipcode string `json:"zipcode"`
	City    string `json:"city"`
	Street  string `json:"street"`
}

type Media struct {
	Config  map[string]string `json:"config"`
	Posts   []graphql.ID      `json:"posts"`
	Stories []graphql.ID      `json:"stories"`
}

type Enterprise struct {
	ID        graphql.ID  `json:"id" bson:"_id,omitempty"`
	Name      string      `json:"name"`
	Logo      string      `json:"logo"`
	Address   *Address    `json:"address"`
	Employees *[]Employee `json:"employees"`
	Media     *Media      `json:"media"`
}

type EnterpriseQuery struct{}

func (*EnterpriseQuery) Enterprise(ctx context.Context, args struct{ Index int }) (*Enterprise, error) {
	dbctx, cancel := createDBContext()
	defer cancel()

	// fetch the user
	user := &usergql{}

	fetchErr := database.UserCol.FindOne(dbctx, bson.M{"_id": ctx.Value(User_id).(primitive.ObjectID)}).Decode(&user)
	if fetchErr != nil {
		return nil, fetchErr
	}

	// get the id of the enterprise and fetch it
	enterprise := &Enterprise{}
	finderr := database.EnterpriseCol.FindOne(dbctx, bson.M{"_id": (*user.ENTERPRISES)[args.Index]}).Decode(enterprise)
	if finderr != nil {
		return nil, finderr
	}
	return enterprise, nil
}

// Enterprise mutations
type EnterpriseMutation struct{}

// create a new enterprise provided that it has a unique name
func (*EnterpriseMutation) AddEnterprise(ctx context.Context, args struct{ Data Enterprise }) (*Enterprise, error) {
	dbctx, cancel := createDBContext()
	defer cancel()

	// chek id the name is unique
	if count, _ := database.EnterpriseCol.CountDocuments(ctx, bson.M{"name": args.Data.Name}); count > 0 {
		return nil, errors.New("an enterprise with the same name already exists")
	}

	// insert
	insertresult, dberr := database.EnterpriseCol.InsertOne(dbctx, args.Data)
	if dberr != nil {
		return nil, dberr
	}

	// find the inserted enterprise
	enterprise := &Enterprise{}

	finderr := database.EnterpriseCol.FindOne(ctx, bson.M{"_id": insertresult.InsertedID}).Decode(enterprise)
	if finderr != nil {
		return nil, finderr
	}

	// update the user's enterprises list
	_, updateErr := database.UserCol.UpdateByID(dbctx, ctx.Value(User_id).(primitive.ObjectID), bson.M{"$addToSet": bson.M{"enterprises": enterprise.ID}})
	if updateErr != nil {
		return nil, updateErr
	}

	// return the result
	return enterprise, nil
}

type message struct {
	Message string `json:"message"`
}

func (*EnterpriseMutation) DeleteEnterprise(ctx context.Context, args struct {
	Password     string
	EnterpriseID graphql.ID
}) (*message, error) {
	dbctx, cancel := createDBContext()
	defer cancel()

	// check if the enterprise exists
	if count, _ := database.EnterpriseCol.CountDocuments(ctx, bson.M{"_id": args.EnterpriseID}); count <= 0 {
		return nil, errors.New("an enterprise with the same name already exists")
	}
	// compare passwords
	fetchuser := auth.User{}
	database.UserCol.FindOne(ctx, bson.M{"_id": ctx.Value(User_id).(primitive.ObjectID)}).Decode(&fetchuser)
	passerr := bcrypt.CompareHashAndPassword([]byte(fetchuser.Password), []byte(args.Password))
	if passerr != nil {
		return nil, passerr
	}

	// find the enterprise
	enterprise := &Enterprise{}

	finderr := database.EnterpriseCol.FindOne(ctx, bson.M{"_id": args.EnterpriseID}).Decode(enterprise)
	if finderr != nil {
		return nil, finderr
	}

	// extract employees' refs
	employees := make([]graphql.ID, len(*enterprise.Employees))
	for i, e := range *enterprise.Employees {
		employees[i] = e.Ref
	}
	// delete the enterprise entry on other Users' documents
	database.UserCol.UpdateMany(dbctx, bson.M{"_id": bson.M{"$in": employees}}, bson.M{"$pull": bson.M{"enterprises": args.EnterpriseID}})

	// delete the enterprise
	_, deleteerr := database.EnterpriseCol.DeleteOne(dbctx, bson.M{"_id": args.EnterpriseID})
	if deleteerr != nil {
		return nil, deleteerr
	}
	return &message{"enterprise successfully deleted"}, nil
}
