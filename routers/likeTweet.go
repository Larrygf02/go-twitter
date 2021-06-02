package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-twitter/bd"
	"github.com/go-twitter/models"
)

func RegisterLikeTweet(w http.ResponseWriter, r *http.Request) {
	var like models.TweetLike
	err := json.NewDecoder(r.Body).Decode(&like)

	if err != nil {
		http.Error(w, "Ocurrio en el cuerpo del tweet", http.StatusBadRequest)
		return
	}

	register := models.TweetLike{
		UserId:      IdUser,
		CreatedDate: time.Now(),
		Tweet:       like.Tweet,
	}

	_, err = bd.AddLikeTweet(register)
	if err != nil {
		http.Error(w, "Ocurrio un error al momento de grabar el like", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func GetLikeTweet(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID ", http.StatusBadRequest)
		return
	}

	response, status := bd.GetLikesTweet(ID)
	if !status {
		http.Error(w, "Ocurrio un error al momento de obtener los likes", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
