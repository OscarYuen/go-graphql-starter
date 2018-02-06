package service

import (
	h "github.com/OscarYuen/go-graphql-starter/handler"
	"github.com/OscarYuen/go-graphql-starter/model"
	"github.com/spf13/viper"
	"log"
	"testing"
	"time"
)

var (
	authService         *AuthService
	appName             string
	signedSecret        string
	expiredTimeInSecond time.Duration
	debugMode           bool
	logFormat           string
)

func init() {
	viper.SetConfigName("Config")
	viper.AddConfigPath("../")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Unable to connect to db: %s \n", err)
	}
	log := h.NewLogger(&appName, debugMode, &logFormat)
	appName = viper.Get("app-name").(string)
	signedSecret = viper.Get("auth.jwt-secret").(string)
	debugMode = viper.Get("log.debug-mode").(bool)
	logFormat = viper.Get("log.log-format").(string)
	expiredTimeInSecond = time.Duration(viper.Get("auth.jwt-expire-in").(int64))
	authService = NewAuthService(&appName, &signedSecret, &expiredTimeInSecond, log)
}

func TestSignJWT(t *testing.T) {
	user := &model.User{
		ID:       "1",
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
