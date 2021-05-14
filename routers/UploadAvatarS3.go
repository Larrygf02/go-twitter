package routers

import (
	"net/http"
	"strings"

	"github.com/go-twitter/bd"
	"github.com/go-twitter/models"
)

func UploadAvatarS3(w http.ResponseWriter, r *http.Request) {
	file, handler, _ := r.FormFile("avatar")
	var extension = strings.Split(handler.Filename, ".")[1]
	var filelocation string = "images/avatars/" + IdUser + "." + extension
	var s3 = models.S3
	var bucket string = "apps-go-twittor"
	s3 = new(models.S3Client)
	s3.Region = "us-east-1"
	s3.NewSession(s3.Region)
	_, err := s3.Upload(file, bucket, filelocation)
	if err != nil {
		http.Error(w, "Error al cargar la imagen!"+err.Error(), http.StatusBadRequest)
		return
	}
	// Update Avatar
	var user models.User
	var status bool
	user.Avatar = "https://" + bucket + ".s3.amazonaws.com/" + filelocation
	status, err = bd.EditUser(user, IdUser)
	if err != nil || !status {
		http.Error(w, "Error al grabar el avatar en el user", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}
