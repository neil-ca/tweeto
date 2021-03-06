package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/Neil-uli/tweeto/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ReadAllUsers(ID string, page int64, search string, typ string) ([]*models.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoC.Database("tewto")
	col := db.Collection("users")

	var results []*models.User

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	// M is an unordered representation of a BSON document. bson.M{"foo": "bar", "hello": "world", "pi": 3.14159}
	query := bson.M{
		"name": bson.M{"$regex": `(?i)` + search},
	}
	cur, err := col.Find(ctx, query, findOptions)
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}
	var found, include bool

	for cur.Next(ctx) {
		var s models.User
		err := cur.Decode(&s)
		if err != nil {
			fmt.Println(err.Error())
			return results, false
		}

		var r models.Relation
		r.UserID = ID
		r.UserRelationID = s.ID.Hex()

		include = false

		found, err = ConsultRelation(r)
		if typ == "new" && found == false {
			include = true
		}
		if typ == "follow" && found == true {
			include = true
		}

		if r.UserRelationID == ID {
			include = false
		}
		if include == true {
			findOptions.SetProjection(
				bson.M{
					"email":     0,
					"password":  0,
					"banner":    0,
					"biography": 0,
					"location":  0,
					"siteWeb":   0,
				},
			)
			results = append(results, &s)
		}
	}
	err = cur.Err()
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}
	cur.Close(ctx)
	return results, true
}
