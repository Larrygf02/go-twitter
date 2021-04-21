package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/go-twitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetProfile(ID string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("twitter")
	collection := db.Collection("usuarios")

	var profile models.User
	objectId, _ := primitive.ObjectIDFromHex(ID) // nos devuelve un objeto de tipo ObjectID
	condition := bson.M{
		"_id": objectId,
	}
	err := collection.FindOne(ctx, condition).Decode(&profile)
	profile.Password = "" // no debemos pasar el pass en el objeto
	if err != nil {
		fmt.Println("Registro no encontrado " + err.Error())
		return profile, err
	}
	return profile, nil

}
