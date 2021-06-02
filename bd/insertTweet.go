package bd

import (
	"context"
	"time"

	"github.com/go-twitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertTweet(tweet models.Tweet) (string, bool, error) {
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
