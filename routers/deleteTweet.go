package routers

import (
	"net/http"

	"github.com/go-twitter/bd"
)

func DeleteTweet(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID ", http.StatusBadRequest)
		return
	}

	err := bd.BorroTweet(ID, IdUser)
	if err != nil {
		http.Error(w, "Ocurrió un error al momento de guardar el tweet", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
