package cmd

import (
	"github.com/amirhnajafiz/go-graphql/internal/cmd/server"
	"github.com/amirhnajafiz/go-graphql/internal/gql"
)

func Execute() {
	s := gql.Init()
	app := server.Init(s)
	_ = app.SetTrustedProxies([]string{"192.168.1.2"})

	err := app.Run(":5000")
	if err != nil {
		panic(err)
	}
}
