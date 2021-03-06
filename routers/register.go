package routers

import (
	"encoding/json"
	"net/http"

	"github.com/go-twitter/bd"
	"github.com/go-twitter/models"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user) // el body es un stream (una vez se lee)
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}
	if len(user.Email) == 0 {
		http.Error(w, "El email del usuario es requerido", 400)
		return
	}

	if len(user.Password) < 6 {
		http.Error(w, "El password es requerido", 400)
		return
	}

	_, found, _ := bd.ExistUser(user.Email)
	if found {
		http.Error(w, "Ya existe un usuario registrado con ese email", 400)
		return
	}

	_, status, err := bd.InsertUser(user)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar registrar el usuario", 400)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado insertar el registro del usuario", 400)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	var response models.ResponseRelation
	response.Status = true
	json.NewEncoder(w).Encode(response)
}
