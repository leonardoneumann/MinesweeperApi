package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/leonardoneumann/minesweeperapi/api/database"
	"github.com/leonardoneumann/minesweeperapi/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const ContentType string = "Content-Type"
const ContentTypeValue string = "application/json; charset=UTF-8"

// GetGame : Gets an existing game
func GetGame(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, ContentTypeValue)

	collection := database.Connect("games")

	// we get params with mux.
	var params = mux.Vars(r)

	id, _ := primitive.ObjectIDFromHex(params["id"])

	filter := bson.M{"_id": id}

	var game models.Game
	err := collection.FindOne(context.TODO(), filter).Decode(&game)

	if err != nil {
		database.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(game)
	w.WriteHeader(http.StatusOK)
}

// ListGames : List all games
func ListGames(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, ContentTypeValue)

	collection := database.Connect("games")
	var games []models.Game

	// bson.M{},  we passed empty filter. So we want to get all data.
	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		database.GetError(err, w)
		return
	}

	// Close the cursor once finished
	/*A defer statement defers the execution of a function until the surrounding function returns.
	simply, run cur.Close() process but after cur.Next() finished.*/
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var game models.Game
		// & character returns the memory address of the following variable.
		err := cur.Decode(&game) // decode similar to deserialize process.
		if err != nil {
			log.Fatal(err)
		}

		// add item our array
		games = append(games, game)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(games) // encode similar to serialize process.

	w.WriteHeader(http.StatusOK)
}

// NewGame : Creates a new Game
func NewGame(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, ContentTypeValue)

	collection := database.Connect("games")

	var game models.Game
	_ = json.NewDecoder(r.Body).Decode(&game)

	result, err := collection.InsertOne(context.TODO(), game)

	if err != nil {
		database.GetError(err, w)
		return
	}

	insertedId := result.InsertedID
	game.Id = insertedId.(primitive.ObjectID).Hex()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(game)
}

// Play : Makes a play on a given cell
func Play(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, ContentTypeValue)
	w.WriteHeader(http.StatusOK)
}
