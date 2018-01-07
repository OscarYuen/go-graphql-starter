package handler

import (
	"../model"
	"../service"
	"../config"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/net/context"
	"log"
	"net"
	"net/http"
	"strings"
	"encoding/base64"
	"errors"
)

func Authenticate(ctx context.Context, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var isAuthorized bool = false
		tokens, ok := r.Header["Authorization"]
		if ok && len(tokens) >= 1 {
			tokenString := tokens[0]
			tokenString = strings.TrimPrefix(tokenString, "Bearer ")
			token, err := ctx.Value("authService").(*service.AuthService).ValidateJWT(&tokenString)
			if err == nil {
				isAuthorized = true
				if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
					fmt.Println(claims["id"], claims["exp"])
				} else {
					fmt.Println(err)
				}
			}
		}

		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			log.Println(w, "Requester ip: %q is not IP:port", r.RemoteAddr)
		}
		ctx := context.WithValue(ctx, "is_authorized", isAuthorized)
		ctx = context.WithValue(ctx, "requester_ip", ip)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

func Login(ctx context.Context) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		loginResponse  := model.NewLoginResponse()

		//err := json.NewDecoder(r.Body).Decode(&userCredentials)
		userCredentials, err := validateBasicAuthHeader(r)
		if  err != nil {
			//loginResponse.Response.Code = http.StatusBadRequest
			//loginResponse.Response.Message = err.Error()
			http.Error(w, err.Error() ,http.StatusBadRequest)
			return
		}
		result, user := ctx.Value("userService").(*service.UserService).ComparePassword(userCredentials)
		if !result {
			//loginResponse.Response.Code = http.StatusUnauthorized
			//loginResponse.Response.Message = "Unauthorized"
			http.Error(w, config.UnauthorizedAccess ,http.StatusUnauthorized)
			return
		}

		tokenString, err := ctx.Value("authService").(*service.AuthService).SignJWT(user)
		if err != nil {
			http.Error(w, config.TokenError ,http.StatusBadRequest)
			return
			//loginResponse.Response.Code = http.StatusBadRequest
			//loginResponse.Response.Message = "Sign Error"
		}
		loginResponse.Response.Code = http.StatusOK
		loginResponse.JWT = *tokenString
		jsonResponse, _ := json.Marshal(loginResponse)
		w.WriteHeader(loginResponse.Response.Code)
		w.Write(jsonResponse)
	})
}
func validateBasicAuthHeader(r *http.Request) (*model.UserCredentials, error) {
	auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	if len(auth) != 2 || auth[0] != "Basic" {
		return nil, errors.New(config.CredentialsError)
	}
	payload, _ := base64.StdEncoding.DecodeString(auth[1])
	pair := strings.SplitN(string(payload), ":", 2)
	if len(pair) != 2 {
		return nil, errors.New(config.CredentialsError)
	}
	userCredentials := model.UserCredentials{
		Email:    pair[0],
		Password: pair[1],
	}
	return &userCredentials, nil
}
