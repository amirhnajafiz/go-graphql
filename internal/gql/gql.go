package gql

import (
	"fmt"
	"log"

	"github.com/amirhnajafiz/go-graphql/internal/store"
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func Init() graphql.Schema {
	db, err := gorm.Open("sqlite3", "tutorials.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	db.AutoMigrate(&store.Author{})
	db.AutoMigrate(&store.Book{})

	aggregateSchema := graphql.Fields{
		"author": store.SetupSingleAuthorSchema(),
	}
	aggregateMutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"create": store.SetupAuthorMutations(),
			"add":    store.SetupBookMutation(),
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
