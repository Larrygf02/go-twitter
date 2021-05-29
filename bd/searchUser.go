package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/go-twitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SearchUser(ID string, page int64, search string, type_user string) ([]*models.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	collection := db.Collection("usuarios")

	var results []*models.User

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	query := bson.M{
		"name": bson.M{"$regex": `(?i)` + search},
	}
	cursor, err := collection.Find(ctx, query, findOptions)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Error 1")
		return results, false
	}
	var found, include bool
	for cursor.Next(ctx) {
		var filterUser models.User
		err := cursor.Decode(&filterUser)
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println("Error 2")
			return results, false
		}
		var relation models.Relation
		relation.UserId = ID
		relation.UserRelation = filterUser.ID.Hex()

		include = false
		found, _ = ExistRelation(relation)
		// usuarios a los que no sigo
		if type_user == "new" && !found {
			include = true
		}

		// usuarios a los que sigo
		if type_user == "follow" && found {
			include = true
		}

		if relation.UserRelation == ID {
			include = false
		}

		if include {
			filterUser.Password = ""
			filterUser.Biography = ""
			filterUser.Website = ""
			filterUser.Location = ""
			filterUser.Banner = ""
			filterUser.Email = ""
			results = append(results, &filterUser)
		}
	}

	err = cursor.Err()
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	cursor.Close(ctx)
	return results, true
}
