package bd

import (
	"context"
	"time"

	"github.com/go-twitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func EditUser(user models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	collection := db.Collection("usuarios")

	register := make(map[string]interface{})
	if len(user.Name) > 0 {
		register["name"] = user.Name
	}

	if len(user.Surname) > 0 {
		register["surname"] = user.Surname
	}
	register["birthdate"] = user.BirthDate

	if len(user.Banner) > 0 {
		register["banner"] = user.Banner
	}

	if len(user.Biography) > 0 {
		register["biography"] = user.Biography
	}

	if len(user.Location) > 0 {
		register["location"] = user.Location
	}

	if len(user.Website) > 0 {
		register["website"] = user.Website
	}

	if len(user.Avatar) > 0 {
		register["avatar"] = user.Avatar
	}

	updateQuery := bson.M{
		"$set": register,
	}

	objId, _ := primitive.ObjectIDFromHex(ID)
	filterByID := bson.M{
		"_id": bson.M{"$eq": objId},
	}

	_, err := collection.UpdateOne(ctx, filterByID, updateQuery)
	if err != nil {
		return false, err
	}
	return true, nil
}
