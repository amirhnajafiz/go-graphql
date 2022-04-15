package gql

import (
	"log"

	"github.com/graphql-go/graphql"
)

func Init() graphql.Schema {
	fields := graphql.Fields{
		"Authors": &graphql.Field{
			Type: AuthorType(),
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery), Mutation: AuthorMutation()}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	return schema
}
