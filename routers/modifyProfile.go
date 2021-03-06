package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Neil-uli/tweeto/bd"
	"github.com/Neil-uli/tweeto/models"
)

func ModifyProfile(w http.ResponseWriter, r *http.Request) {
	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Incorrect data "+err.Error(), 400)
		return
	}
	var status bool
	status, err = bd.ModifyRegister(t, IDUser)
	if err != nil {
		http.Error(w, "Error trying to edit the registry"+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "It was not modified", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
