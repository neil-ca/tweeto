package routers

import (
	"net/http"

	"github.com/ulicod3/tweeto/bd"
	"github.com/ulicod3/tweeto/models"
)

func DownRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	var t models.Relation
	t.UserID = IDUser
	t.UserRelationID = ID

	status, err := bd.DeleteRelation(t)
	if err != nil {
		http.Error(w, "error deleting relation"+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "could not delete relationship"+err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
