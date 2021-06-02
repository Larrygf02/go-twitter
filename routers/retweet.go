package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-twitter/bd"
	"github.com/go-twitter/models"
)

func RegisterRetweet(w http.ResponseWriter, r *http.Request) {
	var retweet models.Retweet
	err := json.NewDecoder(r.Body).Decode(&retweet)

	if err != nil {
		http.Error(w, "Ocurrio en el cuerpo del retweet", http.StatusBadRequest)
		return
	}

	register := models.Retweet{
		UserId:      IdUser,
		CreatedDate: time.Now(),
		Tweet:       retweet.Tweet,
	}

	_, err = bd.AddRetweet(register)
	if err != nil {
		http.Error(w, "Ocurrio un error al momento de grabar el like", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func GetRetweets(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID ", http.StatusBadRequest)
		return
	}

	response, status := bd.GetRetweets(ID)
	if !status {
		http.Error(w, "Ocurrio un error al momento de obtener los likes", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
