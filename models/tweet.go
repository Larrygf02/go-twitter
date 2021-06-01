package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tweet struct {
	UserId         string    `bson:"userid" json:"userid,omitempty"`
	Message        string    `bson:"message" json:"message,omitempty"`
	CreatedDate    time.Time `bson:"date" json:"date,omitempty"`
	IsComment      bool      `bson:"is_comment" json:"is_comment, omitempty"`
	TwitterComment string    `bson:"twitter_comment" json:"twitter_comment, omitempty"`
	IsRetweet      bool      `bson:"is_retweet" json:"is_retweet, omitempty"`
	TwitterRetweet string    `bson:"twitter_retweet" json:"twitter_retweet, omitempty"`
}

type GetTweet struct {
	ID             primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserId         string             `bson:"userid" json:"userid,omitempty"`
	Message        string             `bson:"message" json:"message,omitempty"`
	CreatedDate    time.Time          `bson:"date" json:"date,omitempty"`
	IsComment      bool               `bson:"is_comment" json:"is_comment, omitempty"`
	TwitterComment string             `bson:"twitter_comment" json:"twitter_comment, omitempty"`
	IsRetweet      bool               `bson:"is_retweet" json:"is_retweet, omitempty"`
	TwitterRetweet string             `bson:"twitter_retweet" json:"twitter_retweet, omitempty"`
}
