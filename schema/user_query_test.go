package schema

import (
	"testing"

	"../conf"
	"github.com/OscarYuen/go-graphql-example/schema"
	"github.com/neelance/graphql-go"
	"github.com/neelance/graphql-go/gqltesting"
)

var rootSchema = graphql.MustParseSchema(schema.GetRootSchema(), &schema.Resolver{})
var db = conf.ConnectDB("../test.db")

func init() {
	schema.SetDatabase(db)
}

func TestBasic(t *testing.T) {
	gqltesting.RunTests(t, []*gqltesting.Test{
		{
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
