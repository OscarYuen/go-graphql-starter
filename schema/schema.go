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

func SetDatabase(db *gorm.DB) {
	DB = db
}

type Resolver struct{}
