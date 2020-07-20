package routers

import (
	"github.com/Neil-uli/tewto/bd"
	"github.com/Neil-uli/tewto/models"
	"io"
	"net/http"
	"os"
	"strings"
)

func UpBanner(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("banner")
	var extension = strings.Split(handler.Filename,".")[1]
	var archive string = "uploads/banners/" +IDUser + "." + extension

	f, err := os.OpenFile(archive, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "banner upload error"+err.Error(),http.StatusBadRequest)
		return
	}
	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "banner copy error"+err.Error(),http.StatusBadRequest)
		return
	}

	var user models.User
	var status bool

	user.Banner = IDUser + "." + extension
	status, err = bd.ModifyRegister(user, IDUser)
	if err != nil || status == false {
		http.Error(w, "Error saving banner in bd"+err.Error(),http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
