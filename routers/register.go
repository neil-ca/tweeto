package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Neil-uli/tewto/bd"
	"github.com/Neil-uli/tewto/models"
)

/*Register create in bd*/
func Register(w http.ResponseWriter, r *http.Request) {

	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Email is required", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "Min 6 characters required", 400)
		return
	}

	_, found, _ := bd.CheckIfExistUser(t.Email)
	if found == true {
		http.Error(w, "One user already exists with this email", 400)
		return
	}

	_, status, err := bd.InsertRegistry(t)
	if err != nil {
		http.Error(w, "Failed to register a user"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "unable to insert user", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
