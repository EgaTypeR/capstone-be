package utils

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

const DatabaseName = "CrimeAlertCapstone"

func ConnectDB() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(os.Getenv("DB_URL")) // Replace with your MongoDB URI
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	db := client.Database(DatabaseName)
	DB = db

	log.Println("Connected to MongoDB!")
	return client, nil
}
