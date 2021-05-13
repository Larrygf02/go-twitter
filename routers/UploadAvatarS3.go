package routers

import (
	"net/http"
	"strings"

	"github.com/go-twitter/models"
)

func UploadAvatarS3(w http.ResponseWriter, r *http.Request) {
	file, handler, _ := r.FormFile("avatar")
	var extension = strings.Split(handler.Filename, ".")[1]
	var filelocation string = "images/avatars/" + IdUser + "." + extension
	var s3 = models.S3
	s3 = new(models.S3Client)
	s3.Region = "us-east-1"
	s3.NewSession(s3.Region)
	_, err := s3.Upload(file, "apps-go-twittor", filelocation)
	if err != nil {
		http.Error(w, "Error al cargar la imagen!"+err.Error(), http.StatusBadRequest)
		return
	}

}
