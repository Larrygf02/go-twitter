package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-twitter/bd"
)

func ReadTweetsFollowers(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Debe enviar el parametro pagina", http.StatusBadRequest)
		return
	}
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "Debe enviar el parametro pagina", http.StatusBadRequest)
		return
	}

	results, status := bd.ReadTweetsFollowers(IdUser, page)
	if !status {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(results)

}
