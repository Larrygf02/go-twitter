package bd

import (
	"context"
	"time"

	"github.com/go-twitter/models"
	"go.mongodb.org/mongo-driver/bson"
)

func ExistUser(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	collection := db.Collection("usuarios")
	condition := bson.M{"email": email}
	var result models.User

	err := collection.FindOne(ctx, condition).Decode(&result)
	ID := result.ID.Hex()
	if err != nil {
		return result, false, ID
	}
	return result, true, ID
}
