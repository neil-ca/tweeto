package routers

import (
	"github.com/Neil-uli/tewto/bd"
	"net/http"
)

func DeleteTweet(w http.ResponseWriter, r *http.Request){
	ID := r.URL.Query().Get("id")
	if len(ID)<1 {
		http.Error(w, "I need id",http.StatusBadRequest)
		return
	}

	err := bd.DeleteTweet(ID, IDUser)
	if err != nil {
		http.Error(w, "I cant delete"+err.Error(),http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
