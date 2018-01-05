package service

import (
	"../model"
	"encoding/base64"
	"github.com/dgrijalva/jwt-go"
	"sync"
	"time"
	"fmt"
	"strings"
	"log"
)

var (
	authServiceInstance *AuthService
	authOnce            sync.Once
)

type AuthService struct {
}

func NewAuthService() *AuthService {
	authOnce.Do(func() {
		authServiceInstance = &AuthService{}
	})
	return authServiceInstance
}

func (a *AuthService) SignJWT(user *model.User) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":         base64.StdEncoding.EncodeToString([]byte(string(user.ID))),
		"created_at": user.CreatedAt,
		"exp":        time.Now().Add(time.Second * 400).Unix(),
		"iss":        "go-grapql-starter",
	})

	tokenString, err := token.SignedString([]byte("1234"))
	return &tokenString, err
}

func (a *AuthService) ValidateJWT(tokens []string) (*jwt.Token, error){
	tokenString := tokens[0]
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("1234"), nil
	})
	log.Println(err)
	return token, err
}
