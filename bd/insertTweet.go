package bd

import (
	"context"
	"time"
	"github.com/Neil-uli/tewto/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertTweet(t models.RecordTweet)(string, bool, error){
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoC.Database("tewto")
	col := db.Collection("tweet")

	registry := bson.M{
		"userid": t.UserID,
		"message": t.Message,
		"date": t.Date,
	}
	result, err := col.InsertOne(ctx, registry)
	if err != nil {
		return "",false,err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}
