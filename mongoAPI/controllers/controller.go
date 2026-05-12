package controllers

// import "github.com/kishanghosh090/mongoAPI/db"
import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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

func getAllMovie() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var movies []primitive.M

	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, primitive.M(movie))
	}
	defer cursor.Close(context.Background())
	return movies
}

// actual controllers

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allMovies := getAllMovie()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie models.Netflix

	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params)
}
