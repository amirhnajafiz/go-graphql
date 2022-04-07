package gql

import (
	"github.com/graphql-go/graphql"
)

func AuthorType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "author",
		Fields: graphql.Fields{
			"Name": &graphql.Field{
				Type: graphql.String,
			},
			"Books": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
		},
	})
}
