package controllers

// import "github.com/kishanghosh090/mongoAPI/db"
import (
	"context"
	"fmt"
	"log"

	"github.com/kishanghosh090/mongoAPI/db"
	"github.com/kishanghosh090/mongoAPI/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
)

var collection = db.Collection

// mongodb helpers

func insertOneMovie(movie models.Netflix) {
	inserted, err := collection.InsertOne(context.TODO(), movie)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("inserted one movie id : ", inserted.InsertedID)
}

// update 1

func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	res, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	println("modified count ", res.ModifiedCount)
}

func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)

	filter := bson.M{"_id": id}

	res, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		log.Fatal(err)
	}

	println("modified count ", res.DeletedCount)
}
