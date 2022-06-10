package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ulicod3/tweeto/bd"
)

func ReadTweets(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "You must send the id", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "You must send the page", http.StatusBadRequest)
		return
	}
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "You must send the page with a value greater than 0", http.StatusBadRequest)
		return
	}
	pag := int64(page)
	response, correct := bd.ReadTweets(ID, pag)
	if correct == false {
		http.Error(w, "Error reading the tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
