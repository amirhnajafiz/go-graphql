package gql

import (
	"fmt"

	"github.com/amirhnajafiz/go-graphql/internal/store"
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func Init(cfg Config) (graphql.Schema, error) {
	db, err := gorm.Open("sqlite3", cfg.Database)
	if err != nil {
		return graphql.Schema{}, err
	}

	defer db.Close()

	db.AutoMigrate(&store.Author{})
	db.AutoMigrate(&store.Book{})

	aggregateSchema := graphql.Fields{
		"author": store.SetupSingleAuthorSchema(),
		"books":  store.SetupSingleBookSchema(),
	}
	aggregateMutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"create": store.CreateAuthorMutations(cfg.Database),
			"add":    store.AddBookMutation(cfg.Database),
		},
	})

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: aggregateSchema}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery), Mutation: aggregateMutation}

	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		return graphql.Schema{}, err
	}

	return schema, nil
}

func ExecuteQuery(query string, schema graphql.Schema) (*graphql.Result, error) {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	if len(result.Errors) > 0 {
		return nil, fmt.Errorf("wrong result, unexpected errors: %v", result.Errors)
	}

	return result, nil
}
