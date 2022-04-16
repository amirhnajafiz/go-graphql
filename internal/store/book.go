package store

import (
	"time"

	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
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

func AddBookMutation(database string) *graphql.Field {
	return &graphql.Field{
		Name: "add",
		Type: GetBookType(),
		Args: graphql.FieldConfigArgument{},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			book := Book{
				Title:       p.Args["title"].(string),
				Reference:   p.Args["title"].(string),
				PublishDate: p.Args["title"].(time.Time),
				Sells:       p.Args["title"].(int),
				Price:       p.Args["title"].(float64),
			}

			db, _ := gorm.Open("sqlite3", database)
			db.Save(&book)

			return book, nil
		},
	}
}
