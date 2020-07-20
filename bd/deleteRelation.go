package bd

import (
	"context"
	"github.com/Neil-uli/tewto/models"
	"time"
)

func DeleteRelation(t models.Relation)(bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoC.Database("tewto")
	col := db.Collection("relation")

	_, err := col.DeleteOne(ctx, t)
	if err != nil{
		return false, err
	}
	return true, nil
}
