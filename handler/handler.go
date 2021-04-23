package handler

import (
	"log"
	"net/http"
	"os"

	"github.com/go-twitter/middlew"
	"github.com/go-twitter/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func HandlerRouters() {
	router := mux.NewRouter()
	router.HandleFunc("/register", middlew.CheckBD(routers.Register)).Methods("POST")
	router.HandleFunc("/login", middlew.CheckBD(routers.Login)).Methods("POST")
	router.HandleFunc("/view-profile", middlew.CheckBD(middlew.ValidateJWT(routers.ViewProfile))).Methods("GET")
	router.HandleFunc("/edit-profile", middlew.CheckBD(middlew.ValidateJWT(routers.EditProfile))).Methods("PUT")
	router.HandleFunc("/tweet", middlew.CheckBD(middlew.ValidateJWT(routers.RegisterTweet))).Methods("POST")
	router.HandleFunc("/tweet", middlew.CheckBD(middlew.ValidateJWT(routers.DeleteTweet))).Methods("DELETE")
	router.HandleFunc("/tweets", middlew.CheckBD(middlew.ValidateJWT(routers.GetTweets))).Methods("GET")
	router.HandleFunc("/upload-avatar", middlew.CheckBD(middlew.ValidateJWT(routers.UploadAvatar))).Methods("POST")
	router.HandleFunc("/avatar", middlew.CheckBD(routers.GetAvatar)).Methods("GET") // no es necesario token
	router.HandleFunc("/upload-banner", middlew.CheckBD(middlew.ValidateJWT(routers.UploadBanner))).Methods("POST")
	router.HandleFunc("/banner", middlew.CheckBD(routers.GetBanner)).Methods("GET") // no es necesario token
	router.HandleFunc("/relation", middlew.CheckBD(middlew.ValidateJWT(routers.InsertRelation))).Methods("POST")
	router.HandleFunc("/relation", middlew.CheckBD(middlew.ValidateJWT(routers.DeleteRelation))).Methods("DELETE")
	router.HandleFunc("/relation", middlew.CheckBD(middlew.ValidateJWT(routers.ExistRelation))).Methods("GET")
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8081"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
