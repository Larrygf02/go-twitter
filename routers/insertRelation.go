package routers

import (
	"net/http"

	"github.com/go-twitter/bd"
	"github.com/go-twitter/models"
)

func InsertRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "El parametro ID es obligatorio", http.StatusBadRequest)
		return
	}
	var relation models.Relation
	relation.UserId = IdUser
	relation.UserRelation = ID

	status, err := bd.InsertRelation(relation)
	if err != nil {
		http.Error(w, "Ocurrio al insertar relacion", http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado insertar la relacion", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)

}
