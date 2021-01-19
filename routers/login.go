package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Neil-uli/tewto/bd"
	"github.com/Neil-uli/tewto/jwt"
	"github.com/Neil-uli/tewto/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Invalid username and password"+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "The email field is required"+err.Error(), 400)
		return
	}
	document, exist := bd.TryLogin(t.Email, t.Password)
	if exist == false {
		http.Error(w, "Invalid username and password", 400)
		return
	}
	jwtKey, err := jwt.GenerateJWT(document)
	if err != nil {
		http.Error(w, "an error occurred while trying to generate the corresponding token"+err.Error(), 400)
		return
	}
	resp := models.ResponseLogin{
		Token: jwtKey,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
