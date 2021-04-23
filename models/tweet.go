package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tweet struct {
	UserId      string    `bson:"userid" json:"userid,omitempty"`
	Message     string    `bson:"message" json:"message,omitempty"`
	CreatedDate time.Time `bson:"date" json:"date,omitempty"`
}

type GetTweet struct {
	ID          primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserId      string             `bson:"userid" json:"userid,omitempty"`
	Message     string             `bson:"message" json:"message,omitempty"`
	CreatedDate time.Time          `bson:"date" json:"date,omitempty"`
}
