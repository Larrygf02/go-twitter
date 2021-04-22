package routers

import (
	"encoding/json"
	"net/http"

	"github.com/go-twitter/bd"
)

func ViewProfile(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}
	profile, err := bd.GetProfile(ID)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar el buscar el registro "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(profile)
}
