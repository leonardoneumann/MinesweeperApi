package api

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/leonardoneumann/minesweeperapi/api/database"
	"github.com/leonardoneumann/minesweeperapi/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Authenticate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	collection := database.Connect("users")

	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)

	filter := bson.M{"username": user.Username, "password": user.Password}

	err := collection.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
		helper.GetError(err, w)
		return
	} else {
		var token models.AuthToken
		token.Timestamp = time.Now()
		tokens := database.Connect("authTokens")
		result, err := collection.InsertOne(context.TODO(), token)

		if err != nil {
			database.GetError(err, w)
			return
		}

		insertedId := result.InsertedID
		token.Id = insertedId.(primitive.ObjectID).Hex()

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
	}
}
