package handler

import (
	"../model"
	"../service"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/net/context"
	"log"
	"net"
	"net/http"
)

func Authenticate(ctx context.Context, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokens, ok := r.Header["Authorization"]
		if ok && len(tokens) >= 1 {
			token, err := ctx.Value("authService").(*service.AuthService).ValidateJWT(tokens)
			if err == nil {
				if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
					fmt.Println(claims["id"], claims["exp"])
				} else {
					fmt.Println(err)
				}
			}
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

func Login(ctx context.Context) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var (
			userCredentials model.UserCredentials
			loginResponse   = model.NewLoginResponse()
		)

		err := json.NewDecoder(r.Body).Decode(&userCredentials)
		if err != nil {
			log.Println(err)
			loginResponse.Response.Code = http.StatusBadRequest
			loginResponse.Response.Message = "Invalid Format"
		}
		result, user := ctx.Value("userService").(*service.UserService).ComparePassword(&userCredentials)
		if !result {
			loginResponse.Response.Code = http.StatusUnauthorized
			loginResponse.Response.Message = "Unauthorized"
		} else {
			tokenString, err := ctx.Value("authService").(*service.AuthService).SignJWT(user)
			if err != nil {
				log.Println(err)
				loginResponse.Response.Code = http.StatusBadRequest
				loginResponse.Response.Message = "Sign Error"
			}
			loginResponse.Response.Code = http.StatusOK
			loginResponse.JWT = *tokenString
		}
		jsonResponse, _ := json.Marshal(loginResponse)
		w.WriteHeader(loginResponse.Response.Code)
		w.Write(jsonResponse)
	})
}
