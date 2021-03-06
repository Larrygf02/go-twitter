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
	router.HandleFunc("/tweet", middlew.CheckBD(middlew.ValidateJWT(routers.GetTweet))).Methods("GET")
	router.HandleFunc("/tweets", middlew.CheckBD(middlew.ValidateJWT(routers.GetTweets))).Methods("GET")
	router.HandleFunc("/upload-avatar", middlew.CheckBD(middlew.ValidateJWT(routers.UploadAvatarS3))).Methods("POST")
	router.HandleFunc("/avatar", middlew.CheckBD(routers.GetAvatar)).Methods("GET") // no es necesario token
	router.HandleFunc("/upload-banner", middlew.CheckBD(middlew.ValidateJWT(routers.UploadBannerS3))).Methods("POST")
	router.HandleFunc("/banner", middlew.CheckBD(routers.GetBanner)).Methods("GET") // no es necesario token
	router.HandleFunc("/relation", middlew.CheckBD(middlew.ValidateJWT(routers.InsertRelation))).Methods("POST")
	router.HandleFunc("/relation", middlew.CheckBD(middlew.ValidateJWT(routers.DeleteRelation))).Methods("DELETE")
	router.HandleFunc("/relation", middlew.CheckBD(middlew.ValidateJWT(routers.ExistRelation))).Methods("GET")
	router.HandleFunc("/list-user", middlew.CheckBD(middlew.ValidateJWT(routers.ViewUsers))).Methods("GET")
	router.HandleFunc("/list-tweets", middlew.CheckBD(middlew.ValidateJWT(routers.ReadTweetsFollowers))).Methods("GET") // leer tweets
	router.HandleFunc("/following", middlew.CheckBD(middlew.ValidateJWT(routers.GetFollowing))).Methods("GET")
	router.HandleFunc("/followers", middlew.CheckBD(middlew.ValidateJWT(routers.GetFollowers))).Methods("GET")
	router.HandleFunc("/tweet/comment", middlew.CheckBD(middlew.ValidateJWT(routers.RegisterCommentTweet))).Methods("POST")
	router.HandleFunc("/tweet/comments", middlew.CheckBD(middlew.ValidateJWT(routers.GetCommentsTweet))).Methods("GET")
	router.HandleFunc("/tweet/quotetweet", middlew.CheckBD(middlew.ValidateJWT(routers.RegisterQuoteTweet))).Methods("POST")
	router.HandleFunc("/tweet/quotetweets", middlew.CheckBD(middlew.ValidateJWT(routers.GetQuotesTweet))).Methods("GET")
	router.HandleFunc("/tweet/like", middlew.CheckBD(middlew.ValidateJWT(routers.RegisterLikeTweet))).Methods("POST")
	router.HandleFunc("/tweet/likes", middlew.CheckBD(middlew.ValidateJWT(routers.GetLikeTweet))).Methods("GET")
	router.HandleFunc("/tweet/dislike", middlew.CheckBD(middlew.ValidateJWT(routers.DislikeTweet))).Methods("DELETE")
	router.HandleFunc("/tweet/retweet", middlew.CheckBD(middlew.ValidateJWT(routers.RegisterRetweet))).Methods("POST")
	router.HandleFunc("/tweet/retweets", middlew.CheckBD(middlew.ValidateJWT(routers.GetRetweets))).Methods("GET")
	router.HandleFunc("/tweet/unretweet", middlew.CheckBD(middlew.ValidateJWT(routers.UnRetweet))).Methods("DELETE")
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8081"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
