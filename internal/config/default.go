package config

import (
	"github.com/amirhnajafiz/go-graphql/internal/gql"
	"github.com/amirhnajafiz/go-graphql/internal/logger"
)

func Default() Config {
	return Config{
		Port:  ":5000",
		Proxy: "0.0.0.0",
		GQL: gql.Config{
			Database: "authors.db",
		},
		Logger: logger.Config{
			Level: "debug",
		},
	}
}
