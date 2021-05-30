package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/go-twitter/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetFollowing(ID string) ([]*models.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	collection := db.Collection("relacion")
	condition := bson.M{"userid": ID}
	idUsers := make([]string, 1)
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
		idUsers = append(idUsers, relation.UserRelation)
	}
	fmt.Println(idUsers)
	return results, true
}
