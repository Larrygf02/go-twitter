package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/go-twitter/models"
	"go.mongodb.org/mongo-driver/bson"
)

// Obtener quienes comentaron el tweet

func GetCommentsTweet(ID string) ([]*models.Tweet, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	collection := db.Collection("tweet")
	condition := bson.M{"twitter_comment": ID}
	var results []*models.Tweet

	cursor, _ := collection.Find(ctx, condition)
	if err := cursor.All(ctx, &results); err != nil {
		fmt.Println(err.Error())
		return results, err
	}
	return results, nil
}
