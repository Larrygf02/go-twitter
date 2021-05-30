package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/go-twitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetFollowing(ID string) ([]*models.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	collection := db.Collection("relacion")
	condition := bson.M{"userid": ID}
	idUsers := make([]primitive.ObjectID, 0)
	var results []*models.User

	cursor, err := collection.Find(ctx, condition)
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	for cursor.Next(ctx) {
		var relation models.Relation
		err := cursor.Decode(&relation)
		if err != nil {
			fmt.Println(err.Error())
			return results, false
		}
		objectId, _ := primitive.ObjectIDFromHex(relation.UserRelation)
		idUsers = append(idUsers, objectId)
	}
	collection = db.Collection("usuarios")
	cursor, _ = collection.Find(ctx, bson.M{"_id": bson.M{"$in": idUsers}})
	if err = cursor.All(ctx, &results); err != nil {
		fmt.Println(err.Error())
		return results, false
	}
	fmt.Println(idUsers)
	return results, true
}

func GetFollowers(ID string) ([]*models.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	collection := db.Collection("relacion")
	condition := bson.M{"userrelation": ID}
	idUsers := make([]primitive.ObjectID, 0)
	var results []*models.User

	cursor, err := collection.Find(ctx, condition)
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	for cursor.Next(ctx) {
		var relation models.Relation
		err := cursor.Decode(&relation)
		if err != nil {
			fmt.Println(err.Error())
			return results, false
		}
		objectId, _ := primitive.ObjectIDFromHex(relation.UserRelation)
		idUsers = append(idUsers, objectId)
	}
	collection = db.Collection("usuarios")
	cursor, _ = collection.Find(ctx, bson.M{"_id": bson.M{"$in": idUsers}})
	if err = cursor.All(ctx, &results); err != nil {
		fmt.Println(err.Error())
		return results, false
	}
	fmt.Println(idUsers)
	return results, true
}
