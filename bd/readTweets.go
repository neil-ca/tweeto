package bd

import (
	"context"
	"log"
	"time"

	"github.com/ulicod3/tweeto/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ReadTweets(ID string, page int64) ([]*models.GetTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoC.Database("tewto")
	col := db.Collection("tweet")

	var results []*models.GetTweets

	condition := bson.M{
		"userid": ID,
	}
	options := options.Find()
	options.SetLimit(20)
	options.SetSort(bson.D{{Key: "date", Value: -1}})
	options.SetSkip((page - 1) * 20)

	cursor, err := col.Find(ctx, condition, options)
	if err != nil {
		log.Fatal(err.Error())
		return results, false
	}
	for cursor.Next(context.TODO()) {
		var registry models.GetTweets
		err := cursor.Decode(&registry)
		if err != nil {
			return results, false
		}
		results = append(results, &registry)
	}
	return results, true

}
