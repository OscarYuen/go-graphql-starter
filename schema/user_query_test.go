package schema

import (
	"testing"

	"../conf"
	"github.com/neelance/graphql-go"
	"github.com/neelance/graphql-go/gqltesting"
)

var rootSchema = graphql.MustParseSchema(GetRootSchema(), &Resolver{})

func init() {
	conf.ConnectDB("../test.db")
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
