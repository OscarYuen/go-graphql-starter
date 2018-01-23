package main

import (
	"github.com/OscarYuen/go-graphql-starter/config"
	h "github.com/OscarYuen/go-graphql-starter/handler"
	"github.com/OscarYuen/go-graphql-starter/model"
	"github.com/OscarYuen/go-graphql-starter/resolver"
	"github.com/OscarYuen/go-graphql-starter/schema"
	"github.com/OscarYuen/go-graphql-starter/service"
	"log"
	"net/http"

	"github.com/OscarYuen/go-graphql-starter/loader"
	graphql "github.com/neelance/graphql-go"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"time"
)

func main() {
	db, err := config.OpenDB("test.db")
	if err != nil {
		log.Fatalf("Unable to connect to db: %s \n", err)
	}
	viper.SetConfigName("Config")
	viper.AddConfigPath(".")
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error config file: %s \n", err)
	}

	var (
		appName             = viper.Get("app-name").(string)
		signedSecret        = viper.Get("auth.jwt-secret").(string)
		expiredTimeInSecond = time.Duration(viper.Get("auth.jwt-expire-in").(int64))
		debugMode           = viper.Get("log.debug_mode").(bool)
		logFormat           = viper.Get("log.log_format").(string)
	)
	notificationHub := model.NewNotificationHub()
	go notificationHub.Run()
	ctx := context.WithValue(context.Background(), "userService", service.NewUserService(db))
	ctx = context.WithValue(ctx, "authService", service.NewAuthService(&appName, &signedSecret, &expiredTimeInSecond))
	ctx = context.WithValue(ctx, "notificationHub", notificationHub)

	graphqlSchema := graphql.MustParseSchema(schema.GetRootSchema(), &resolver.Resolver{})

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "graphiql.html")
	}))

	http.Handle("/login", h.AddContext(ctx, h.Login()))

	http.Handle("/ws", h.AddContext(ctx, h.Authenticate(h.WebSocket(notificationHub))))

	http.HandleFunc("/notification", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "notification.html")
	}))

	logger := &h.Logger{&appName, debugMode, &logFormat}
	http.Handle("/query", h.AddContext(ctx, logger.Logging(h.Authenticate(&h.GraphQL{Schema: graphqlSchema, Loaders:loader.NewLoaderCollection()}))))

	log.Fatal(http.ListenAndServe(":3000", nil))
}
