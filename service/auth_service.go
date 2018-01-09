package service

import (
	"../model"
	"encoding/base64"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"sync"
	"time"
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
		"id":         base64.StdEncoding.EncodeToString([]byte(strconv.FormatInt(user.ID, 10))),
		"created_at": user.CreatedAt,
		"exp":        time.Now().Add(time.Second * 400).Unix(),
		"iss":        "go-grapql-starter",
	})

	tokenString, err := token.SignedString([]byte("1234"))
	return &tokenString, err
}

func (a *AuthService) ValidateJWT(tokenString *string) (*jwt.Token, error) {
	token, err := jwt.Parse(*tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("1234"), nil
	})
	return token, err
}
