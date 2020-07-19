package bd

import (
	"context"
	"time"
	"fmt"
	"github.com/Neil-uli/tewto/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
func SearchProfile(ID string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoC.Database("tewto")
	col := db.Collection("users")

	var profile models.User
	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id":objID,
	}

	err:=col.FindOne(ctx, condition).Decode(&profile)
	profile.Password=""
	if err != nil {
		fmt.Println("Record not found"+err.Error())
		return profile, err
	}
	return profile, nil
}
