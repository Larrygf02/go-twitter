package jwtoken

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-twitter/models"
)

func GenerateToken(user models.User) (string, error) {
	key_code := []byte("GOLANGTWITTER")
	payload := jwt.MapClaims{
		"email":      user.Email,
		"name":       user.Name,
		"surname":    user.Surname,
		"birth_date": user.BirthDate,
		"location":   user.Location,
		"website":    user.Website,
		"avatar":     user.Avatar,
		"_id":        user.ID.Hex(),
		"expire":     time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(key_code)
	if err != nil {
		fmt.Println(err)
		return tokenStr, err
	}
	return tokenStr, nil
}
