package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const connectionString = "mongodb://admin:1234@localhost:27017"
const dbName = "netflix"
const colName = "watchlist"

var Collection *mongo.Collection

// connect with mongodb

func init() {
	// client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// connect to MONGODB (v2 mongo.Connect takes variadic *options.ClientOptions)
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// verify connection
	if err = client.Ping(context.TODO(), nil); err != nil {
		log.Fatal(err)
	}

	// set collection
	Collection = client.Database(dbName).Collection(colName)

	// collection instance ready
}
