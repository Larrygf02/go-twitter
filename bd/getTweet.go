package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/go-twitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetTweet(ID string, page int64) ([]models.GetTweet, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	collection := db.Collection("tweet")

	var results []models.GetTweet
	conditions := make([]bson.M, 0)
	conditions = append(conditions, bson.M{"$match": bson.M{"userid": ID}})
	conditions = append(conditions, bson.M{
		"$project": bson.M{
			"userid": bson.M{
				"$toObjectId": "$userid",
			},
			"message":         1,
			"date":            1,
			"is_comment":      1,
			"is_retweet":      1,
			"twitter_comment": 1,
			"twitter_retweet": 1,
		},
	})
	conditions = append(conditions, bson.M{
		"$lookup": bson.M{
			"from":         "usuarios",
			"localField":   "userid",
			"foreignField": "_id",
			"as":           "userid",
		},
	})
	conditions = append(conditions, bson.M{"$unwind": "$userid"})
	conditions = append(conditions, bson.M{"$sort": bson.M{
		"date": -1,
	}})
	skip := (page - 1) * 8
	conditions = append(conditions, bson.M{"$skip": skip})
	conditions = append(conditions, bson.M{"$limit": 8})
	cursor, _ := collection.Aggregate(ctx, conditions)
	err := cursor.All(ctx, &results)
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}
	fmt.Println(results)
	return results, true
}

func GetOneTweet(ID string) (models.GetTweet, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	collection := db.Collection("tweet")
	objectId, _ := primitive.ObjectIDFromHex(ID)

	var result []models.GetTweet

	conditions := make([]bson.M, 0)
	conditions = append(conditions, bson.M{"$match": bson.M{"_id": objectId}})
	conditions = append(conditions, bson.M{
		"$project": bson.M{
			"userid": bson.M{
				"$toObjectId": "$userid",
			},
			"message":         1,
			"date":            1,
			"is_comment":      1,
			"is_retweet":      1,
			"twitter_comment": 1,
			"twitter_retweet": 1,
		},
	})
	conditions = append(conditions, bson.M{
		"$lookup": bson.M{
			"from":         "usuarios",
			"localField":   "userid",
			"foreignField": "_id",
			"as":           "userid",
		},
	})
	conditions = append(conditions, bson.M{"$unwind": "$userid"})
	//condition := bson.M{
	//	"_id": objectId,
	//}

	//err := collection.FindOne(ctx, conditions).Decode(&result)
	cursor, _ := collection.Aggregate(ctx, conditions)
	err := cursor.All(ctx, &result)
	if err != nil {
		fmt.Println("Registro no encontrado " + err.Error())
		return result[0], err
	}
	return result[0], nil
}
