package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoC obj*/
var MongoC = ConnectionBD()

var clientOptions = options.Client().ApplyURI("mongodb+srv://dbUser:zPoCaCD9tEl5CyX2@cluster0.1sokj.mongodb.net/tweeto?retryWrites=true&w=majority")

/*ConnectionBD func for conn*/
func ConnectionBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("successful connection !")
	return client
}

/*CheckConnection up*/
func CheckConnection() int {
	err := MongoC.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
