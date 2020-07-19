package bd

import (
	"context"
	"time"
	"github.com/Neil-uli/tewto/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ModifyRegister(u models.User, ID string) (bool,error){
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoC.Database(("tewto"))
	col := db.Collection("users")

	registry := make(map[string]interface{})
	if len(u.Name) > 0 {
		registry["name"] = u.Name
	}
	if len(u.Surname) > 0 {
		registry["surname"] = u.Surname
	}
	registry["dateofbirth"] = u.DateOfBirth

	if len(u.Avatar) > 0 {
		registry["avatar"] = u.Avatar
	}
	if len(u.Banner) > 0 {
		registry["banner"] = u.Banner
	}
	if len(u.Biography) > 0 {
		registry["biography"] = u.Biography
	}
	if len(u.Location) > 0 {
		registry["location"] = u.Location
	}
	if len(u.SiteWeb) > 0 {
		registry["siteweb"] = u.SiteWeb
	}

	updtString := bson.M{
		"$set": registry,
	}
	objID, _:= primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id":bson.M{"$eq":objID}}

	_, err := col.UpdateOne(ctx, filter, updtString)
	if err != nil {
		return false, err
	}
	return true, nil
}
