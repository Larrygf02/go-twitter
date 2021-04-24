package bd

import (
	"context"
	"time"

	"github.com/go-twitter/models"
	"go.mongodb.org/mongo-driver/bson"
)

func ReadTweetsFollowers(ID string, page int) ([]models.ReadTweetsFollowers, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	collection := db.Collection("relacion")

	skip := (page - 1) * 20
	conditions := make([]bson.M, 0)
	conditions = append(conditions, bson.M{"$match": bson.M{"userid": ID}})
	conditions = append(conditions, bson.M{
		"$lookup": bson.M{
			"from":         "tweet",
			"localField":   "userrelation",
			"foreignField": "userid",
			"as":           "tweet",
		}})
	conditions = append(conditions, bson.M{"$unwind": "$tweet"})
	conditions = append(conditions, bson.M{"$sort": bson.M{
		"date": -1,
	}})
	conditions = append(conditions, bson.M{"$skip": skip})
	conditions = append(conditions, bson.M{"$limit": 20})

	cursor, _ := collection.Aggregate(ctx, conditions)
	var result []models.ReadTweetsFollowers
	err := cursor.All(ctx, &result)
	if err != nil {
		return result, false
	}
	return result, true
}
