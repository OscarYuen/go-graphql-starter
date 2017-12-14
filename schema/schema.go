package schema

import (
	"github.com/jinzhu/gorm"
	"bytes"
)

var DB *gorm.DB
var Schema = `
	schema {
		query: Query
		mutation: Mutation
	}
	type Query {
		findUserByEmail(email: String!): User
	}
	type Mutation {
		createUser(email: String!, password: String!): User
	}`

func GetRootSchema() string {
	var buffer bytes.Buffer
	buffer.WriteString(Schema)
	buffer.WriteString(userSchema)
	return buffer.String()
}

func SetDatabase(db *gorm.DB) {
	DB = db
}

type Resolver struct{}
