package main

import (
	"./config"
	"./handler"
	"./model"
	"./resolver"
	"./schema"
	"./service"
	"log"
	"net/http"

	"github.com/neelance/graphql-go"
	"github.com/neelance/graphql-go/relay"
	"golang.org/x/net/context"
)

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/home" {
		http.Error(w, "Not found", 404)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	http.ServeFile(w, r, "home.html")
}

func main() {
	db, err := config.OpenDB("test.db")
	if err != nil {
		log.Fatal("Unable to connect to db:")
		log.Fatal(err)
	}
	notificationHub := model.NewNotificationHub()
	go notificationHub.Run()
	ctx := context.WithValue(context.Background(), "userService", service.NewUserService(db))
	ctx = context.WithValue(ctx, "authService", service.NewAuthService())
	ctx = context.WithValue(ctx, "notificationHub", notificationHub)

	graphqlSchema := graphql.MustParseSchema(schema.GetRootSchema(), &resolver.Resolver{})

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(page)
	}))

	http.Handle("/login", handler.Login(ctx))

	http.Handle("/ws", handler.Authenticate(ctx, handler.WebSocket(notificationHub)))

	http.HandleFunc("/home", serveHome)

	http.Handle("/query", handler.Authenticate(ctx, &relay.Handler{Schema: graphqlSchema}))

	log.Fatal(http.ListenAndServe(":3000", nil))
}

var page = []byte(`
<!DOCTYPE html>
<html>
	<head>
		<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.10.2/graphiql.css" />
		<script src="https://cdnjs.cloudflare.com/ajax/libs/fetch/1.1.0/fetch.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/react/15.5.4/react.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/react/15.5.4/react-dom.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.10.2/graphiql.js"></script>
	</head>
	<body style="width: 100%; height: 100%; margin: 0; overflow: hidden;">
		<div id="graphiql" style="height: 100vh;">Loading...</div>
		<script>
			function graphQLFetcher(graphQLParams) {
				return fetch("/query", {
					method: "post",
					headers: {
					 'Authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjoiMjAxOC0wMS0wM1QxNToyNToyNloiLCJleHAiOjE1MTU1NzMwODAsImlkIjoiTXc9PSIsImlzcyI6ImdvLWdyYXBxbC1zdGFydGVyIn0.c1QYtflkF0ZxQIj05mOf_5BfniZ9ePDXajapKge-EQg',
				   	},
					body: JSON.stringify(graphQLParams),
					credentials: "include",
				}).then(function (response) {
					return response.text();
				}).then(function (responseBody) {
					try {
						return JSON.parse(responseBody);
					} catch (error) {
						return responseBody;
					}
				});
			}
			ReactDOM.render(
				React.createElement(GraphiQL, {fetcher: graphQLFetcher}),
				document.getElementById("graphiql")
			);
		</script>
	</body>
</html>
`)
