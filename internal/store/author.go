package store

import (
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

type Author struct {
	gorm.Model
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
				Type: graphql.NewList(GetBookType()),
			},
		},
	},
)

func GetAuthorType() *graphql.Object {
	return authorType
}

func SetupSingleAuthorSchema() *graphql.Field {
	return &graphql.Field{
		Type: GetAuthorType(),
	}
}

func SetupAuthorMutations() *graphql.Field {
	return &graphql.Field{
		Name: "create",
		Type: GetAuthorType(),
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			author := Author{
				Name: p.Args["name"].(string),
			}

			db, _ := gorm.Open("sqlite3", "authors.db")
			db.Save(&author)

			return author, nil
		},
	}
}
