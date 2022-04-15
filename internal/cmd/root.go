package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/amirhnajafiz/go-graphql/internal/gql"
)

func Execute() {
	s := gql.Init()
	query := `
		{
			author {
				Name
			}
		}
	`
	insert := `
		mutation {
			create(name: "Amir") {
				name
			}
		}
	`

	_ = gql.ExecuteQuery(insert, s)
	r := gql.ExecuteQuery(query, s)
	rJSON, _ := json.Marshal(r)

	fmt.Printf("%s \n", rJSON)
}
