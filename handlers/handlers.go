package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/ulicod3/tweeto/middlew"
	"github.com/ulicod3/tweeto/routers"
)

func Handlers() {
	router := mux.NewRouter()
	router.HandleFunc("/register", middlew.CheckBD(routers.Register)).Methods("POST")
	router.HandleFunc("/login", middlew.CheckBD(routers.Login)).Methods("POST")
	router.HandleFunc("/view/profile", middlew.CheckBD(middlew.ValidateJWT(routers.ViewProfile))).Methods("GET")
	router.HandleFunc("/profile", middlew.CheckBD(middlew.ValidateJWT(routers.ModifyProfile))).Methods("PUT")
	router.HandleFunc("/tweet", middlew.CheckBD(middlew.ValidateJWT(routers.RecordTweet))).Methods("POST")
	router.HandleFunc("/tweets", middlew.CheckBD(middlew.ValidateJWT(routers.ReadTweets))).Methods("GET")
	router.HandleFunc("/tweet", middlew.CheckBD(middlew.ValidateJWT(routers.DeleteTweet))).Methods("DELETE")

	router.HandleFunc("/avatar-up", middlew.CheckBD(middlew.ValidateJWT(routers.UpAvatar))).Methods("POST")
	router.HandleFunc("/avatar", middlew.CheckBD(routers.GetAvatar)).Methods("GET")

	router.HandleFunc("/banner-up", middlew.CheckBD(middlew.ValidateJWT(routers.UpBanner))).Methods("POST")
	router.HandleFunc("/banner", middlew.CheckBD(routers.GetBanner)).Methods("GET")
	router.HandleFunc("/follow", middlew.CheckBD(routers.UpRelation)).Methods("POST")
	router.HandleFunc("/unfollow", middlew.CheckBD(routers.DownRelation)).Methods("DELETE")
	router.HandleFunc("/relation", middlew.CheckBD(routers.ConsultRelation)).Methods("GET")
	router.HandleFunc("/list/users", middlew.CheckBD(routers.ListUsers)).Methods("GET")
	router.HandleFunc("/tweetsFollowers", middlew.CheckBD(routers.ReadTweetsFollowers)).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
