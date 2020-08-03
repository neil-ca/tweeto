package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/Neil-uli/tewto/middlew"
	"github.com/Neil-uli/tewto/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Handlers() {
	router := mux.NewRouter()
	router.HandleFunc("/register", middlew.CheckBD(routers.Register)).Methods("POST")
	router.HandleFunc("/login", middlew.CheckBD(routers.Login)).Methods("POST")
	router.HandleFunc("/view-profile", middlew.CheckBD(middlew.ValidateJWT(routers.ViewProfile))).Methods("GET")
	router.HandleFunc("/modify-profile", middlew.CheckBD(middlew.ValidateJWT(routers.ModifyProfile))).Methods("PUT")
	router.HandleFunc("/tweet", middlew.CheckBD(middlew.ValidateJWT(routers.RecordTweet))).Methods("POST")
	router.HandleFunc("/read-tweets", middlew.CheckBD(middlew.ValidateJWT(routers.ReadTweets))).Methods("GET")
	router.HandleFunc("/delete-tweet", middlew.CheckBD(middlew.ValidateJWT(routers.DeleteTweet))).Methods("DELETE")

	router.HandleFunc("/up-avatar", middlew.CheckBD(middlew.ValidateJWT(routers.UpAvatar))).Methods("POST")
	router.HandleFunc("/get-avatar", middlew.CheckBD(routers.GetAvatar)).Methods("GET")

	router.HandleFunc("/up-banner", middlew.CheckBD(middlew.ValidateJWT(routers.UpBanner))).Methods("POST")
	router.HandleFunc("/get-banner", middlew.CheckBD(routers.GetBanner)).Methods("GET")
	router.HandleFunc("/up-relation", middlew.CheckBD(routers.UpRelation)).Methods("POST")
	router.HandleFunc("/down-relation", middlew.CheckBD(routers.DownRelation)).Methods("DELETE")
	router.HandleFunc("/consult-relation", middlew.CheckBD(routers.ConsultRelation)).Methods("GET")
	router.HandleFunc("/list-users", middlew.CheckBD(routers.ListUsers)).Methods("GET")
	router.HandleFunc("/read-tweets-followers", middlew.CheckBD(routers.ReadTweetsFollowers)).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
