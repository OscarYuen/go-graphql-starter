package schema

import (
	"bytes"
)

var Schema = `
	schema {
		query: Query
		mutation: Mutation
	}`

func GetRootSchema() string {
	var buffer bytes.Buffer
	buffer.WriteString(Schema)
	buffer.WriteString(`type Query {`)
	buffer.WriteString(userQuery)
	buffer.WriteString(`}`)
	buffer.WriteString(`type Mutation {`)
	buffer.WriteString(userMutation)
	buffer.WriteString(`}`)
	buffer.WriteString(userSchema)
	return buffer.String()
}

type Resolver struct{}
