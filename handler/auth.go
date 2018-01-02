package handler

import (
	"log"
	"net/http"
	"golang.org/x/net/context"
	"net"
	"fmt"
)

func Authenticate(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		log.Print(tokenString)
		if tokenString == "" {
			http.Error(w, "Invalid Token", http.StatusUnauthorized)
			return
		}

		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			fmt.Fprintf(w, "userip: %q is not IP:port", r.RemoteAddr)
		}
		userIP := net.ParseIP(ip)
		if userIP == nil {
			fmt.Fprintf(w, "userip: %q is not IP:port", r.RemoteAddr)
			return
		}
		ctx := context.WithValue(r.Context(), "requester_ip", ip)
		log.Println(ip)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}
