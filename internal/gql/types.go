package gql

import (
	"github.com/amirhnajafiz/go-graphql/internal/store"
	"github.com/graphql-go/graphql"
)

var authors []store.Author

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

func AuthorMutation() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "author",
		Fields: graphql.Fields{
			"create": &graphql.Field{
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					author := store.Author{
						Name: p.Args["name"].(string),
					}

					authors = append(authors, author)

					return author, nil
				},
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
