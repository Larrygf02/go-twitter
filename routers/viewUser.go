package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-twitter/bd"
)

func ViewUsers(w http.ResponseWriter, r *http.Request) {
	typeUser := r.URL.Query().Get("type")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	pagTemp, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "Debe enviar la pagina", http.StatusBadRequest)
		return
	}
	pag := int64(pagTemp)
	result, status := bd.SearchUser(IdUser, pag, search, typeUser)
	if !status {
		http.Error(w, "Error al leer los usuarios", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

}
