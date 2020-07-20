package bd

import (
	"context"
	"fmt"
	"github.com/Neil-uli/tewto/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func ConsultRelation(t models.Relation) (bool,error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoC.Database("tewto")
	col := db.Collection("relation")

	condition := bson.M{
		"userid": t.UserID,
		"userrelationid": t.UserRelationID,
	}

	var result models.Relation
	err := col.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	return true, nil
}
