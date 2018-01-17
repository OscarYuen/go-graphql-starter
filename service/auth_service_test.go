package service

import (
	"github.com/OscarYuen/go-graphql-starter/model"
	"github.com/spf13/viper"
	"testing"
	"time"
	"log"
)

var (
	authService         *AuthService
	appName             string
	signedSecret        string
	expiredTimeInSecond time.Duration
)

func init() {
	viper.SetConfigName("Config")
	viper.AddConfigPath("../")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Unable to connect to db: %s \n", err)
	}
	appName = viper.Get("app-name").(string)
	signedSecret = viper.Get("auth.jwt-secret").(string)
	expiredTimeInSecond = time.Duration(viper.Get("auth.jwt-expire-in").(int64))
	authService = NewAuthService(&appName, &signedSecret, &expiredTimeInSecond)
}

func TestSignJWT(t *testing.T) {
	user := &model.User{
		ID:       1,
		Email:    "test@1.com",
		Password: "123456",
	}
	tokenString, err := authService.SignJWT(user)
	if err != nil {
		t.Errorf("Error during signing JWT")
	}
	if *tokenString == "" || tokenString == nil {
		t.Errorf("Invalid JWT")
	}

}
