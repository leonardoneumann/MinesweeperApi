package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/leonardoneumann/minesweeperapi/api/database"
	"github.com/leonardoneumann/minesweeperapi/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	collection := database.Connect("users")

	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)

	filter := bson.M{"username": user.Username}
	var userExists models.User

	err := collection.FindOne(context.TODO(), filter).Decode(&userExists)

	if err == nil {
		//TODO: should return error message to the client
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("User Exists")
		return
	} else {
		result, err := collection.InsertOne(context.TODO(), user)

		if err != nil {
			database.GetError(err, w)
			return
		}

		insertedId := result.InsertedID
		user.Id = insertedId.(primitive.ObjectID).Hex()

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
	}

}
