package middleware

import (
	"log"
	"net/http"
)

func Preflight(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Globally Handling preflight
		if r.Method == "OPTIONS" {
			log.Println("Handled Preflight")
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
