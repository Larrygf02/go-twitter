package routers

import (
	"errors"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-twitter/bd"
	"github.com/go-twitter/models"
)

var Email string
var IdUser string

func ProcessToken(token string) (*models.Claim, bool, string, error) {
	key_code := []byte("GOLANGTWITTER")
	claims := &models.Claim{}
	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}
	token = strings.TrimSpace(splitToken[1])
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return key_code, nil
	})
	if err == nil {
		_, found, _ := bd.ExistUser(claims.Email)
		if found {
			Email = claims.Email
			IdUser = claims.ID.Hex()
		}
		return claims, found, IdUser, nil
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}
	return claims, false, string(""), err
}
