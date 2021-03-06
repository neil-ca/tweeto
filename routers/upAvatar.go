package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/Neil-uli/tweeto/bd"
	"github.com/Neil-uli/tweeto/models"
)

func UpAvatar(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("avatar")
	var extension = strings.Split(handler.Filename, ".")[1]
	var archive string = "uploads/avatars/" + IDUser + "." + extension

	f, err := os.OpenFile(archive, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "image upload error"+err.Error(), http.StatusBadRequest)
		return
	}
	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "image copy error"+err.Error(), http.StatusBadRequest)
		return
	}

	var user models.User
	var status bool

	user.Avatar = IDUser + "." + extension
	status, err = bd.ModifyRegister(user, IDUser)
	if err != nil || status == false {
		http.Error(w, "Error saving avatar in bd"+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
