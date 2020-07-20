package routers

import (
	"encoding/json"
	"github.com/Neil-uli/tewto/bd"
	"net/http"
	"strconv"
)

func ReadTweetsFollowers(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query().Get("page"))<1 {
		http.Error(w,"I need the param page",http.StatusBadRequest)
		return
	}
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w,"I need the param page > 0",http.StatusBadRequest)
		return
	}
	response, correct := bd.ReadTweetsFollowers(IDUser, page)
	if correct == false {
		http.Error(w, "Error reading tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}