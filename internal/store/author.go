package store

import "github.com/graphql-go/graphql"

type Author struct {
	Name  string
	Books []Book
}

var authorType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Author",
		Fields: graphql.Fields{
			"Name": &graphql.Field{
				Type: graphql.String,
			},
			"Books": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
		},
	},
)
