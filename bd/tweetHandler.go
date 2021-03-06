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
	collection := db.Collection("tweet")
	// condition := bson.M{"twitter_comment": ID}

	conditions := make([]bson.M, 0)
	conditions = append(conditions, bson.M{"$match": bson.M{"twitter_comment": ID}})
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
	var results []*models.GetTweet

	cursor, _ := collection.Aggregate(ctx, conditions)
	err := cursor.All(ctx, &results)
	if err != nil {
		fmt.Println("Registros no encontrados " + err.Error())
		return results, err
	}
	return results, nil
}

// Insert comment in tweets

func InsertCommentTweet(tweet models.Tweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	collection := db.Collection("tweet")

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
	collection := db.Collection("tweet")
	conditions := make([]bson.M, 0)
	conditions = append(conditions, bson.M{"$match": bson.M{"twitter_retweet": ID}})
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
	var results []*models.GetTweet

	cursor, _ := collection.Aggregate(ctx, conditions)
	err := cursor.All(ctx, &results)
	if err != nil {
		fmt.Println("Registros no encontrados " + err.Error())
		return results, err
	}
	return results, nil

	/*cursor, _ := collection.Find(ctx, condition)
	if err := cursor.All(ctx, &results); err != nil {
		fmt.Println(err.Error())
		return results, err
	}
	return results, nil*/
}

// Add like to Tweet

func AddLikeTweet(tweet_like models.TweetLike) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	collection := db.Collection("tweet_like")

	register := bson.M{
		"userid": tweet_like.UserId,
		"date":   tweet_like.CreatedDate,
		"tweet":  tweet_like.Tweet,
	}

	_, err := collection.InsertOne(ctx, register)
	if err != nil {
		return false, err
	}
	return true, nil
}

func GetLikesTweet(ID string) ([]*models.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	collection := db.Collection("tweet_like")
	condition := bson.M{"tweet": ID}
	idUsers := make([]primitive.ObjectID, 0)
	var results []*models.User

	cursor, err := collection.Find(ctx, condition)
	if err != nil {
		return results, false
	}

	for cursor.Next(ctx) {
		var like models.TweetLike
		err := cursor.Decode(&like)
		if err != nil {
			return results, false
		}
		objectId, _ := primitive.ObjectIDFromHex(like.UserId)
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

// Borro like
func DeleteLike(ID string, UserID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	collection := db.Collection("tweet_like")

	condition := bson.M{
		"userid": UserID,
		"tweet":  ID,
	}

	_, err := collection.DeleteOne(ctx, condition)
	return err
}

// Add Retweet

func AddRetweet(retweet models.Retweet) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	collection := db.Collection("tweet_retweet")

	register := bson.M{
		"userid": retweet.UserId,
		"date":   retweet.CreatedDate,
		"tweet":  retweet.Tweet,
	}

	_, err := collection.InsertOne(ctx, register)
	if err != nil {
		return false, err
	}
	return true, nil
}

func GetRetweets(ID string) ([]*models.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	collection := db.Collection("tweet_retweet")
	condition := bson.M{"tweet": ID}
	idUsers := make([]primitive.ObjectID, 0)
	var results []*models.User

	cursor, err := collection.Find(ctx, condition)
	if err != nil {
		return results, false
	}

	for cursor.Next(ctx) {
		var like models.TweetLike
		err := cursor.Decode(&like)
		if err != nil {
			return results, false
		}
		objectId, _ := primitive.ObjectIDFromHex(like.UserId)
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

// UnRetweet
func UnRetweet(ID string, UserID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	collection := db.Collection("tweet_retweet")

	condition := bson.M{
		"userid": UserID,
		"tweet":  ID,
	}

	_, err := collection.DeleteOne(ctx, condition)
	return err
}
