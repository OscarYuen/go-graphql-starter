package main

import (
	"github.com/OscarYuen/go-graphql-starter/config"
	"github.com/OscarYuen/go-graphql-starter/handler"
	"github.com/OscarYuen/go-graphql-starter/model"
	"github.com/OscarYuen/go-graphql-starter/resolver"
	"github.com/OscarYuen/go-graphql-starter/schema"
	"github.com/OscarYuen/go-graphql-starter/service"
	"log"
	"net/http"

	graphql "github.com/neelance/graphql-go"
	relay "github.com/neelance/graphql-go/relay"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"time"
	"github.com/OscarYuen/go-graphql-starter/loader"
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
	)

	notificationHub := model.NewNotificationHub()
	go notificationHub.Run()
	ctx := context.WithValue(context.Background(), "userService", service.NewUserService(db))
	ctx = context.WithValue(ctx, "authService", service.NewAuthService(&appName, &signedSecret, &expiredTimeInSecond))
	ctx = context.WithValue(ctx, "notificationHub", notificationHub)
	ctx = loader.NewLoaderCollection().Attach(ctx)

	graphqlSchema := graphql.MustParseSchema(schema.GetRootSchema(), &resolver.Resolver{})

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "graphiql.html")
	}))

	http.Handle("/login", handler.Login(ctx))

	http.Handle("/ws", handler.Authenticate(ctx, handler.WebSocket(notificationHub)))

	http.HandleFunc("/notification", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "notification.html")
	}))

	http.Handle("/query", handler.Authenticate(ctx, &relay.Handler{Schema: graphqlSchema}))

	log.Fatal(http.ListenAndServe(":3000", nil))
}
