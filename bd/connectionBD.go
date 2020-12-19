package bd

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	//go-lint
	_ "github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _ = godotenv.Load(".env")

/*MongoC obj*/
var MongoC = ConnectionBD()

var clientOptions = options.Client().ApplyURI(os.Getenv("MONGO_URL"))

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
