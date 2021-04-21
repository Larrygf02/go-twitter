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
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8081"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
