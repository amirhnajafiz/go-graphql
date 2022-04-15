package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/amirhnajafiz/go-graphql/internal/gql"
	"github.com/graphql-go/graphql"
)

func Execute() {
	s := gql.Init()
	query := `
		{
			Authors {
				Name
				Books {
					Title
					PublishDate
				}
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

	r := graphql.Do(graphql.Params{
		Schema:        s,
		RequestString: insert,
	})

	params := graphql.Params{Schema: s, RequestString: query}
	r = graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}

	rJSON, _ := json.Marshal(r)
	fmt.Printf("%s \n", rJSON)
}
