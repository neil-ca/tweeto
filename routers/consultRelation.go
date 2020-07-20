package routers

import (
	"encoding/json"
	"github.com/Neil-uli/tewto/bd"
	"github.com/Neil-uli/tewto/models"
	"net/http"
)

func ConsultRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	var t models.Relation
	t.UserID = IDUser
	t.UserRelationID = ID

	var resp models.CunsultRelation

	status, err := bd.ConsultRelation(t)
	if err != nil || status == false {
		resp.Status=false
	}else {
		resp.Status=true
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
