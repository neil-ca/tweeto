package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Neil-uli/tewto/bd"
	"github.com/Neil-uli/tewto/models"
)

func RecordTweet(w http.ResponseWriter, r *http.Request) {
	var message models.Tweet
	err := json.NewDecoder(r.Body).Decode(&message)

	registry := models.RecordTweet{
		UserID: IDUser,
		Message: message.Message,
		Date: time.Now(),
	}

	_, status, err := bd.InsertTweet(registry)
	if err != nil {
		http.Error(w, "error when trying to insert the record" + err.Error(), 400)
		return
	}
	if status == false{
		http.Error(w, "empty register", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
