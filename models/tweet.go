package models

import "time"

type Tweet struct {
	UserId      string    `bson:"userid" json:"userid,omitempty"`
	Message     string    `bson:"message" json:"message,omitempty"`
	CreatedDate time.Time `bson:"date" json:"date,omitempty"`
}
