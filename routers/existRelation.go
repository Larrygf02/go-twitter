package routers

import (
	"encoding/json"
	"net/http"

	"github.com/go-twitter/bd"
	"github.com/go-twitter/models"
)

func existRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	var relation models.Relation
	relation.UserId = IdUser
	relation.UserRelation = ID
	var response models.ResponseRelation

	status, err := bd.ExistRelation(relation)
	if err != nil || !status {
		response.Status = false
	}
	response.Status = true

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
