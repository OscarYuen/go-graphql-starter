package handler

import (
	"golang.org/x/net/context"
	"net/http"
)

func AddContext(ctx context.Context, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}
