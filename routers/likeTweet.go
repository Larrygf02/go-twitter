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
