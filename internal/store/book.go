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
	Author      uint
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

func SetupSingleBookSchema() *graphql.Field {
	return &graphql.Field{
		Type: GetBookType(),
	}
}

func AddBookMutation(database string) *graphql.Field {
	return &graphql.Field{
		Name: "add",
		Type: GetBookType(),
		Args: graphql.FieldConfigArgument{},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			var a Author

			book := Book{
				Title:       p.Args["title"].(string),
				Reference:   p.Args["title"].(string),
				PublishDate: p.Args["title"].(time.Time),
				Sells:       p.Args["title"].(int),
				Price:       p.Args["title"].(float64),
				Author:      p.Args["author_id"].(uint),
			}

			db, _ := gorm.Open("sqlite3", database)

			db.Save(&book)
			db.First(&a, book.Author)

			a.Books = append(a.Books, int(book.ID))

			db.Save(a)

			return book, nil
		},
	}
}
