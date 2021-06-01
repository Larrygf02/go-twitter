package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/go-twitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Obtener quienes comentaron el tweet

func GetCommentsTweet(ID string) ([]*models.GetTweet, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	collection := db.Collection("tweet_beta")
	condition := bson.M{"twitter_comment": ID}
	var results []*models.GetTweet

	cursor, _ := collection.Find(ctx, condition)
	if err := cursor.All(ctx, &results); err != nil {
		fmt.Println(err.Error())
		return results, err
	}
	return results, nil
}

// Insert comment in tweets

func InsertCommentTweet(tweet models.Tweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	collection := db.Collection("tweet_beta")

	register := bson.M{
		"userid":          tweet.UserId,
		"message":         tweet.Message,
		"date":            tweet.CreatedDate,
		"is_comment":      tweet.IsComment,
		"twitter_comment": tweet.TwitterComment,
		"is_retweet":      tweet.IsRetweet,
		"twitter_retweet": tweet.TwitterRetweet,
	}

	result, err := collection.InsertOne(ctx, register)
	if err != nil {
		return string(""), false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}

// Get Quote Twees
func GetQuoteTweet(ID string) ([]*models.GetTweet, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	collection := db.Collection("tweet_beta")
	condition := bson.M{"twitter_retweet": ID}
	var results []*models.GetTweet

	cursor, _ := collection.Find(ctx, condition)
	if err := cursor.All(ctx, &results); err != nil {
		fmt.Println(err.Error())
		return results, err
	}
	return results, nil
}
