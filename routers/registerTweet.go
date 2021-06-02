package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-twitter/bd"
	"github.com/go-twitter/models"
)

func RegisterTweet(w http.ResponseWriter, r *http.Request) {
	var tweet models.Tweet
	err := json.NewDecoder(r.Body).Decode(&tweet)

	if err != nil {
		http.Error(w, "Ocurrio en el cuerpo del tweet", http.StatusBadRequest)
		return
	}

	register := models.Tweet{
		UserId:         IdUser,
		Message:        tweet.Message,
		CreatedDate:    time.Now(),
		IsComment:      false,
		TwitterComment: "",
		IsRetweet:      false,
		TwitterRetweet: "",
	}
	_, status, err := bd.InsertTweet(register)
	if err != nil {
		http.Error(w, "Ocurrio un error al momento de grabar el tweet", http.StatusBadRequest)
		return
	}
	if !status {
		http.Error(w, "Ocurrio un error insesperado", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
