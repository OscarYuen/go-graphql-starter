package handler

import (
	"net/http"
	"log"
)

func Authenticate(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		log.Print(tokenString)
		if tokenString == "" {
			http.Error(w, "Invalid Token", http.StatusUnauthorized)
			return
		}
		h.ServeHTTP(w, r)
	})
}