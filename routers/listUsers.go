package routers

import (
	"encoding/json"
	"github.com/Neil-uli/tewto/bd"
	"net/http"
	"strconv"
)

func ListUsers(w http.ResponseWriter, r *http.Request) {
	typeUser := r.URL.Query().Get("type")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	pagTemp, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "The page parameter must be greater than ",http.StatusBadRequest)
		return
	}

	pag := int64(pagTemp)

	result, status := bd.ReadAllUsers(IDUser, pag, search, typeUser)
	if status == false {
		http.Error(w, "Error to read users ",http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-type", "application/json")

	json.NewEncoder(w).Encode(result)
}
