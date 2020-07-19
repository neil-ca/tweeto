package jwt

import (
	"time"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/Neil-uli/tewto/models"
)

func GenerateJWT(t models.User) (string, error){
	myKey := []byte("mysercretpasswd")

	payload := jwt.MapClaims{
		"email": t.Email,
		"name": t.Name,
		"surname": t.Surname,
		"date_of_birth": t.DateOfBirth,
		"biography": t.Biography,
		"location" : t.Location,
		"site-web": t.SiteWeb,
		"_id": t.ID.Hex(),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(myKey)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr,nil
}
