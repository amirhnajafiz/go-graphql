package store

import (
	"github.com/graphql-go/graphql"
	"time"
)

type Book struct {
	Title       string
	Reference   string
	PublishDate time.Time
	Sells       int
	Price       float64
}

var bookType = graphql.NewObject(graphql.ObjectConfig{
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

func GetBookType() *graphql.Object {
	return bookType
}
