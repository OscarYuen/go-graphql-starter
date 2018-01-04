package handler

import (
	"golang.org/x/net/context"
	"log"
	"net"
	"net/http"
)

func Authenticate(ctx context.Context, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		log.Print(tokenString)
		if tokenString == "" {
			http.Error(w, "Invalid Token", http.StatusUnauthorized)
			return
		}

		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			log.Fatal(w, "Requester ip: %q is not IP:port", r.RemoteAddr)
		}
		ctx := context.WithValue(ctx, "requester_ip", ip)
		log.Println(ip)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}
