package graphql

import (
	"context"
	"errors"
	"path/filepath"

	"github.com/graph-gophers/graphql-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"redinnlabs.com/redinn-core/assets"
	"redinnlabs.com/redinn-core/database"
	"redinnlabs.com/redinn-core/utils"
)

type EnterpriseResolvers struct{}

type GetEnterpriseArgs struct {
	Index float64
}

// query: get an enterprise
func (*EnterpriseResolvers) Enterprise(ctx context.Context, args GetEnterpriseArgs) (*EnterpriseGQL, error) {
	// fetch the user
	user, fetchErr := FindUser(ctx.Value(utils.User_id).(primitive.ObjectID))
	if fetchErr != nil {
		return nil, fetchErr
	}
	// get the id of the enterprise
	id, primerr := primitive.ObjectIDFromHex(string((*user.Enterprises)[int(args.Index)]))
	if primerr != nil {
		return nil, primerr
	}
	// get the id of the enterprise and fetch it
	return FindEnterprise(id)
}

type AddEnterpriseArgs struct {
	Data EnterpriseGQL
}

// mutation: create a new enterprise provided that it has a unique name
func (*EnterpriseResolvers) AddEnterprise(ctx context.Context, args AddEnterpriseArgs) (*EnterpriseGQL, error) {
	dbctx, cancel := utils.CreateDBContext()
	defer cancel()

	var userID = ctx.Value(utils.User_id).(primitive.ObjectID)
	// chek if id the name is unique
	if count, _ := database.EnterpriseCol.CountDocuments(ctx, bson.M{"name": args.Data.Name}); count > 0 {
		return nil, errors.New("an enterprise with the same name already exists")
	}
	// add the user that created the enterprise to Employees
	args.Data.Employees = &[]EmployeeGQL{{Ref: graphql.ID(userID.Hex()), Permissions: "111111"}}
	// insert
	insertresult, dberr := database.EnterpriseCol.InsertOne(dbctx, args.Data)
	if dberr != nil {
		return nil, dberr
	}
	insertID := insertresult.InsertedID.(primitive.ObjectID)
	// find the inserted enterprise
	enterprise, finderr := FindEnterprise(insertID)
	if finderr != nil {
		return nil, finderr
	}
	// create a folder
	folder, foldererr := assets.CreateFolder(assets.ENTERPRISE, insertID.Hex())
	if foldererr != nil {
		database.EnterpriseCol.DeleteOne(dbctx, bson.M{"_id": insertID})
		return nil, foldererr
	}
	args.Data.Folder = folder
	// update the user's enterprises list
	_, updateErr := database.UserCol.UpdateByID(dbctx, userID, bson.M{"$addToSet": bson.M{"enterprises": enterprise.ID}})
	if updateErr != nil {
		return nil, updateErr
	}
	// return the result
	return enterprise, nil
}

type DeleteEnterpriseArgs struct {
	Password     string
	EnterpriseID graphql.ID
}

// mutation: delete an enterprise
func (*EnterpriseResolvers) DeleteEnterprise(ctx context.Context, args DeleteEnterpriseArgs) (*utils.Message, error) {
	dbctx, cancel := utils.CreateDBContext()
	defer cancel()

	ID, primerr := primitive.ObjectIDFromHex(string(args.EnterpriseID))
	if primerr != nil {
		return nil, primerr
	}
	// check if the enterprise exists
	if count, _ := database.EnterpriseCol.CountDocuments(ctx, bson.M{"_id": ID}); count <= 0 {
		return nil, errors.New("an enterprise with the same name already exists")
	}
	// compare passwords
	fetchuser := database.User{}
	database.UserCol.FindOne(ctx, bson.M{"_id": ctx.Value(utils.User_id).(primitive.ObjectID)}).Decode(&fetchuser)
	passerr := bcrypt.CompareHashAndPassword([]byte(fetchuser.Password), []byte(args.Password))
	if passerr != nil {
		return nil, passerr
	}
	// find the enterprise
	enterprise, finderr := FindEnterprise(ID)
	if finderr != nil {
		return nil, finderr
	}
	// extract employees' refs
	employees := make([]primitive.ObjectID, len(*enterprise.Employees))
	// map the employees array objects -> ref strings
	for i, e := range *enterprise.Employees {
		employeeID, _ := primitive.ObjectIDFromHex(string(e.Ref))
		employees[i] = employeeID
	}
	// delete the enterprise entry on other Users' documents
	database.UserCol.UpdateMany(dbctx, bson.M{"_id": bson.M{"$in": employees}}, bson.M{"$pull": bson.M{"enterprises": args.EnterpriseID}})
	// delete the enterprise
	_, deleteerr := database.EnterpriseCol.DeleteOne(dbctx, bson.M{"_id": ID})
	if deleteerr != nil {
		return nil, deleteerr
	}
	return &utils.Message{"enterprise successfully deleted"}, nil
}

// mutation: upload an avatar
// important: also applies to users!
type EnterpriseLogoArgs struct {
	Index    float64
	Filename string
}
type EnterpriseLogoRes struct {
	Uploadtoken string
}

func (*EnterpriseResolvers) EnterpriseLogo(ctx context.Context, args EnterpriseLogoArgs) (*EnterpriseLogoRes, error) {
	dbctx, cancel := utils.CreateDBContext()
	defer cancel()

	hash, uuiderr := assets.GenUUIDv4()
	if uuiderr != nil {
		return nil, uuiderr
	}
	// add the file's extension to the hash
	hash += filepath.Ext(args.Filename)

	user, fetchErr := FindUser(ctx.Value(utils.User_id).(primitive.ObjectID))
	if fetchErr != nil {
		return nil, fetchErr
	}
	// get the id of the enterprise
	id, primerr := primitive.ObjectIDFromHex(string((*user.Enterprises)[int(args.Index)]))
	if primerr != nil {
		return nil, primerr
	}
	// update
	_, updateErr := database.EnterpriseCol.UpdateByID(dbctx, id, bson.M{"$set": bson.M{
		"logo": hash,
	}})
	if updateErr != nil {
		return nil, updateErr
	}
	// create an upload token
	token, tokenErr := assets.CreateUploadToken(hash, assets.PHOTO_UPLOAD_DUR)
	if tokenErr != nil {
		return nil, tokenErr
	}
	return &EnterpriseLogoRes{token}, nil
}
