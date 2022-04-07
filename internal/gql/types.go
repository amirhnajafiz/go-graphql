package gql

import (
	"github.com/graphql-go/graphql"
)

func AuthorType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Author",
		Fields: graphql.Fields{
			"Name": &graphql.Field{
				Type: graphql.String,
			},
			"Books": &graphql.Field{
				Type: graphql.NewList(BookType()),
			},
		},
	})
}

func BookType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Book",
		Fields: graphql.Fields{
			"Title": &graphql.Field{
				Type: graphql.String,
			},
			"Reference": &graphql.Field{
				Type: graphql.String,
			},
			"PublishDate": &graphql.Field{
				Type: graphql.DateTime,
			},
			"Sells": &graphql.Field{
				Type: graphql.Int,
			},
			"Price": &graphql.Field{
				Type: graphql.Float,
			},
		},
	})
}
