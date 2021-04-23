package bd

import (
	"context"
	"time"

	"github.com/go-twitter/models"
)

func DeleteRelation(relation models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	collection := db.Collection("relacion")

	_, err := collection.DeleteOne(ctx, relation)
	if err != nil {
		return false, err
	}
	return true, nil
}
