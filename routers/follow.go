package routers

import (
	"encoding/json"
	"net/http"

	"github.com/go-twitter/bd"
)

func GetFollowing(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}
	users, status := bd.GetFollowing(ID)
	if !status {
		http.Error(w, "Ocurrio un error al encontrar los que estas siguiendo", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(users)

}
