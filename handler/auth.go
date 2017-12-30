package handler

import (
	"log"
	"net/http"
	"golang.org/x/net/context"
)

func Authenticate(ctx context.Context, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		log.Print(tokenString)
		if tokenString == "" {
			http.Error(w, "Invalid Token", http.StatusUnauthorized)
			return
		}
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}
