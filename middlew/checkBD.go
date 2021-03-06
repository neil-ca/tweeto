package middlew

import (
	"net/http"

	"github.com/Neil-uli/tweeto/bd"
)

/*CheckBD middlew*/
func CheckBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.CheckConnection() == 0 {
			http.Error(w, "Connection failed", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
