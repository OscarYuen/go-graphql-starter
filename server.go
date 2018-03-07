package main

import (
	"github.com/OscarYuen/go-graphql-starter/config"
	h "github.com/OscarYuen/go-graphql-starter/handler"
	"github.com/OscarYuen/go-graphql-starter/resolver"
	"github.com/OscarYuen/go-graphql-starter/schema"
	"github.com/OscarYuen/go-graphql-starter/service"
	"log"
	"net/http"

	"github.com/OscarYuen/go-graphql-starter/loader"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"time"
)

func main() {
	viper.SetConfigName("Config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error config file: %s \n", err)
	}

	var (
		appName = viper.Get("app-name").(string)

		host     = viper.Get("db.host").(string)
		port     = viper.Get("db.port").(string)
		user     = viper.Get("db.user").(string)
		password = viper.Get("db.password").(string)
		dbname   = viper.Get("db.dbname").(string)

		signedSecret        = viper.Get("auth.jwt-secret").(string)
		expiredTimeInSecond = time.Duration(viper.Get("auth.jwt-expire-in").(int64))

		debugMode = viper.Get("log.debug-mode").(bool)
		logFormat = viper.Get("log.log-format").(string)
	)
	db, err := config.OpenDB(host, port, user, password, dbname)
	if err != nil {
		log.Fatalf("Unable to connect to db: %s \n", err)
	}
	ctx := context.Background()
	log := h.NewLogger(&appName, debugMode, &logFormat)
	roleService := service.NewRoleService(db, log)
	userService := service.NewUserService(db, roleService, log)
	authService := service.NewAuthService(&appName, &signedSecret, &expiredTimeInSecond, log)

	ctx = context.WithValue(ctx, "debugMode", debugMode)
	ctx = context.WithValue(ctx, "log", log)
	ctx = context.WithValue(ctx, "roleService", roleService)
	ctx = context.WithValue(ctx, "userService", userService)
	ctx = context.WithValue(ctx, "authService", authService)

	graphqlSchema := graphql.MustParseSchema(schema.GetRootSchema(), &resolver.Resolver{})

	http.Handle("/login", h.AddContext(ctx, h.Login()))

	loggerHandler := &h.LoggerHandler{debugMode, log}
	http.Handle("/query", h.AddContext(ctx, loggerHandler.Logging(h.Authenticate(&h.GraphQL{Schema: graphqlSchema, Loaders: loader.NewLoaderCollection()}))))

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "graphiql.html")
	}))

	log.Fatal(http.ListenAndServe(":3000", nil))
}
