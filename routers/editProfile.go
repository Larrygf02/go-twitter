package routers

import (
	"encoding/json"
	"net/http"

	"github.com/go-twitter/bd"
	"github.com/go-twitter/models"
)

func EditProfile(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Datos incorrectos"+err.Error(), http.StatusBadRequest)
		return
	}

	var status bool
	status, err = bd.EditUser(user, IdUser)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar el editar el usuario", http.StatusBadRequest)
		return
	}
	if !status {
		http.Error(w, "No se ha logrado modificar el registro del usuario", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)

}
