package store

import "github.com/graphql-go/graphql"

type Author struct {
	Name  string
	Books []Book
}

var authors []Author

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

func SetupAuthorMutations() graphql.Fields {
	return graphql.Fields{
		"create": &graphql.Field{
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

				authors = append(authors, author)

				return author, nil
			},
		},
	}
}
