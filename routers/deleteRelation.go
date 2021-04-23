package routers

import (
	"net/http"

	"github.com/go-twitter/bd"
	"github.com/go-twitter/models"
)

func DeleteRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	var relation models.Relation
	relation.UserId = IdUser
	relation.UserRelation = ID

	status, err := bd.DeleteRelation(relation)
	if err != nil {
		http.Error(w, "Ocurrio al eliminar relacion", http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado eliminar la relacion", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
