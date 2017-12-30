package schema

import (
	"testing"
	"../model"
	"github.com/neelance/graphql-go"
	"github.com/OscarYuen/graphql-go/gqltesting"
	"golang.org/x/net/context"
	"log"
)

var rootSchema = graphql.MustParseSchema(GetRootSchema(), &Resolver{})
var ctx context.Context

func init() {
	db,err := model.OpenDB("../test.db")
	if err != nil {
		log.Fatal(err)
	}
	ctx = context.WithValue(context.Background(), "db", db)
}

func TestBasic(t *testing.T) {
	gqltesting.RunTests(t, []*gqltesting.Test{
		{
			Context: ctx,
			Schema: rootSchema,
			Query: `
				{
					findUserByEmail(email:"peter1") {
						id
						email
						password
					}
				}
			`,
			ExpectedResult: `
				{
					"findUserByEmail":{
						"email":"peter1",
						"id":"2",
						"password":"test"
					}
				}
			`,
		},
	})
}
