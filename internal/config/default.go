package config

import "github.com/amirhnajafiz/go-graphql/internal/gql"

func Default() Config {
	return Config{
		GQL: gql.Config{
			Database: "authors.db",
		},
	}
}
