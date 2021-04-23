package bd

import (
	"context"
	"log"
	"time"

	"github.com/go-twitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetTweet(ID string, page int64) ([]*models.GetTweet, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	collection := db.Collection("tweet")

	var results []*models.GetTweet
	condition := bson.M{
		"userid": ID,
	}
	options := options.Find()
	options.SetLimit(20)
	options.SetSort(bson.D{
		{Key: "date", Value: -1},
	})
	options.SetSkip((page - 1) * 20)

	cursor, err := collection.Find(ctx, condition, options)
	if err != nil {
		log.Fatal(err.Error())
		return results, false
	}
	for cursor.Next(context.TODO()) {
		var register models.GetTweet
		err := cursor.Decode(&register)
		if err != nil {
			return results, false
		}
		results = append(results, &register)
	}
	return results, true
}
