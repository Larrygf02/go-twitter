package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/go-twitter/bd"
	"github.com/go-twitter/models"
)

func UploadAvatar(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("avatar")
	var extension = strings.Split(handler.Filename, ".")[1]
	var filelocation string = "uploads/avatars/" + IdUser + "." + extension
	f, err := os.OpenFile(filelocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error al subir la imagen!"+err.Error(), http.StatusBadRequest)
		return
	}
	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error al copiar la imagen!"+err.Error(), http.StatusBadRequest)
		return
	}
	var user models.User
	var status bool
	user.Avatar = IdUser + "." + extension
	status, err = bd.EditUser(user, IdUser)
	if err != nil || !status {
		http.Error(w, "Error al grabar el avatar en el user", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}