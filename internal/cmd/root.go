package cmd

import (
	"github.com/amirhnajafiz/go-graphql/internal/cmd/server"
	"github.com/amirhnajafiz/go-graphql/internal/gql"
	"github.com/amirhnajafiz/go-graphql/internal/logger"
	"go.uber.org/zap"
)

func Execute() {
	s := gql.Init()
	l := logger.New(logger.Config{})

	app := server.Server{
		L: l,
		S: s,
	}.Init()
	_ = app.SetTrustedProxies([]string{"0.0.0.0"})

	err := app.Run(":5000")
	if err != nil {
		l.Error("server start failed", zap.Error(err))
	}
}
