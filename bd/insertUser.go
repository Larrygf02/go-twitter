package bd

import (
	"context"
	"time"

	"github.com/go-twitter/models"
	"github.com/go-twitter/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertUser(user models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	collection := db.Collection("usuarios")
	user.Password, _ = utils.CryptPassword(user.Password) // actualizando el password encriptado
	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		return "", false, err
	}
	idUser, _ := result.InsertedID.(primitive.ObjectID)
	return idUser.String(), true, nil
}
