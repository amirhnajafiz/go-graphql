package gql

import (
	"fmt"
	"log"

	"github.com/amirhnajafiz/go-graphql/internal/store"
	"github.com/graphql-go/graphql"
)

func Init() graphql.Schema {
	aggregateSchema := graphql.Fields{
		"author": store.SetupSingleAuthorSchema(),
	}
	aggregateMutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"create": store.SetupAuthorMutations(),
		},
	})

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: aggregateSchema}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery), Mutation: aggregateMutation}

	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	return schema
}

func ExecuteQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}

	return result
}
