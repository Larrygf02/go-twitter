package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-twitter/bd"
	"github.com/go-twitter/models"
)

func RegisterQuoteTweet(w http.ResponseWriter, r *http.Request) {
	var tweet models.GetTweet
	err := json.NewDecoder(r.Body).Decode(&tweet)

	if err != nil {
		http.Error(w, "Ocurrio un error en el cuerpo del tweet", http.StatusBadRequest)
		return
	}

	register := models.Tweet{
		UserId:         IdUser,
		Message:        tweet.Message,
		CreatedDate:    time.Now(),
		IsComment:      false,
		TwitterComment: "",
		IsRetweet:      true,
		TwitterRetweet: tweet.TwitterRetweet,
	}

	_, status, err := bd.InsertCommentTweet(register)
	if err != nil {
		http.Error(w, "Ocurrio un error al momento de grabar el quote", http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "Ocurrio un error inesperado", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func GetQuotesTweet(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro id", http.StatusBadRequest)
		return
	}
	response, error := bd.GetQuoteTweet(ID)
	if error != nil {
		http.Error(w, "Error al leer los comentarios del tweet", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

}
