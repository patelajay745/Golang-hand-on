package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mongoapi/model"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// change passowrd first
const connectionString = "mongodb+srv://patelajay745:<password>@cluster0.ajwkw8y.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

const dbName = "netflix"
const colName = "watchlist"

// most IMP
var collection *mongo.Collection

func init() {
	clientOption := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB Connection success")

	collection = client.Database(dbName).Collection(colName)

	fmt.Println("Collection instance is ready")

}

//MongoDB Helpers -file

func insertOneMovie(movie model.Netflix) (moviename string, movieId string) {
	result, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	insertedID := result.InsertedID // Get the inserted ID

	fmt.Println("Inserted 1 movie in db with id:", insertedID)

	return movie.Movie, insertedID.(primitive.ObjectID).Hex()
}

func updateOneMovie(movieId string) (string, error) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		return "", err
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return "", err
	}

	if result.ModifiedCount == 0 {
		return "", fmt.Errorf("no movie with ID %s was updated", movieId)
	}

	movie := getOneMovie(movieId)
	movieName, ok := movie["movie"].(string)
	if !ok {
		return "", fmt.Errorf("unable to retrieve movie name")
	}

	return movieName, nil
}

func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted count:", result)
}

func deleteMovies() {
	result, err := collection.DeleteMany(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted count:", result.DeletedCount)
}

func getOneMovie(movieId string) primitive.M {
	id, _ := primitive.ObjectIDFromHex(movieId)

	var movie bson.M
	err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&movie)
	if err != nil {
		log.Fatal(err)
	}
	return movie
}

func getAllMovies() []primitive.M {

	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var movies []primitive.M

	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}

	defer cur.Close(context.Background())

	return movies
}

// Actual controller -file

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/Json")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/Json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movieName, movieId := insertOneMovie(movie)
	json.NewEncoder(w).Encode(fmt.Sprintf("%s is inserted with id %s", movieName, movieId))

}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/Json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)

	movieName, err := updateOneMovie(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := fmt.Sprintf("%s is marked as watched", movieName)
	json.NewEncoder(w).Encode(response)

}

func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/Json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneMovie(params["id"])

	json.NewEncoder(w).Encode(fmt.Sprintf("%s is deleted", params["id"]))
}

func DeleteAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/Json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	deleteMovies()

	json.NewEncoder(w).Encode("Deleted all movie")
}

func GetOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/Json")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")

	params := mux.Vars(r)
	movie := getOneMovie(params["id"])
	fmt.Println(movie["movie"].(string))
	json.NewEncoder(w).Encode(movie)

}
