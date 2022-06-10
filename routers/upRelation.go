package routers

import (
	"net/http"

	"github.com/ulicod3/tweeto/bd"
	"github.com/ulicod3/tweeto/models"
)

func UpRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "I need ID", http.StatusBadRequest)
		return
	}

	var t models.Relation
	t.UserID = IDUser
	t.UserRelationID = ID

	status, err := bd.InsertRelation(t)
	if err != nil {
		http.Error(w, "error inserting relation"+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "could not insert relationship"+err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
