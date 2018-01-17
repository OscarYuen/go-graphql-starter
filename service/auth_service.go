package service

import (
	"encoding/base64"
	"fmt"
	"github.com/OscarYuen/go-graphql-starter/model"
	jwt "github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

type AuthService struct {
	appName             *string
	signedSecret        *string
	expiredTimeInSecond *time.Duration
}

func NewAuthService(appName *string, signedSecret *string, expiredTimeInSecond *time.Duration) *AuthService {
	return &AuthService{appName, signedSecret, expiredTimeInSecond}
}

func (a *AuthService) SignJWT(user *model.User) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":         base64.StdEncoding.EncodeToString([]byte(strconv.FormatInt(user.ID, 10))),
		"created_at": user.CreatedAt,
		"exp":        time.Now().Add(time.Second * *a.expiredTimeInSecond).Unix(),
		"iss":        *a.appName,
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
