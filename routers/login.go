package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-twitter/bd"
	"github.com/go-twitter/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Usuario o contraseña invalidos "+err.Error(), 400)
		return
	}
	if len(user.Email) == 0 {
		http.Error(w, "Email del usuario es requerido", 400)
		return
	}
	document, exist := bd.Login(user.Email, user.Password)
	if !exist {
		http.Error(w, "Usuario o contraseña invalidos ", 400)
		return
	}
	jwtKey, err := jwt.GenerateToken(document)
	if err != nil {
		http.Error(w, "Ocurrió un error al generar el token "+err.Error(), 400)
		return
	}
	responseLogin := models.ResponseLogin{
		Token: jwtKey,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseLogin)

	// grabar cookie
	expireTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expireTime,
	})
}
