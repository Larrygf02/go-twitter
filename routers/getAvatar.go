package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/go-twitter/bd"
)

func GetAvatar(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	profile, err := bd.GetProfile(ID)
	if err != nil {
		http.Error(w, "Usuario no encontrado", http.StatusBadRequest)
		return
	}
	openFile, err := os.Open("uploads/avatars/" + profile.Avatar)
	if err != nil {
		http.Error(w, "Imagen no encontrada", http.StatusBadRequest)
		return
	}
	_, err = io.Copy(w, openFile)
	if err != nil {
		http.Error(w, "Imagen al copiar", http.StatusBadRequest)
	}
}
