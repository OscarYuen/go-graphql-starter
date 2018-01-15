package service

import (
	"github.com/OscarYuen/go-graphql-starter/model"
	"encoding/base64"
	"fmt"
	"strconv"
	"time"
	jwt "github.com/dgrijalva/jwt-go"
)

type AuthService struct {
	signedSecret *string
	expiredTimeInSecond *time.Duration
}

func NewAuthService(signedSecret *string,expiredTimeInSecond *time.Duration) *AuthService {
	return &AuthService{signedSecret,expiredTimeInSecond}
}

func (a *AuthService) SignJWT(user *model.User) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":         base64.StdEncoding.EncodeToString([]byte(strconv.FormatInt(user.ID, 10))),
		"created_at": user.CreatedAt,
		"exp":        time.Now().Add(time.Second * *a.expiredTimeInSecond).Unix(),
		"iss":        "go-grapql-starter",
	})

	tokenString, err := token.SignedString([]byte(*a.signedSecret))
	return &tokenString, err
}

func (a *AuthService) ValidateJWT(tokenString *string) (*jwt.Token, error) {
	token, err := jwt.Parse(*tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(*a.signedSecret), nil
	})
	return token, err
}
