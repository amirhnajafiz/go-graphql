package cmd

import (
	"github.com/amirhnajafiz/go-graphql/internal/cmd/server"
	"github.com/amirhnajafiz/go-graphql/internal/config"
	"github.com/amirhnajafiz/go-graphql/internal/gql"
	"github.com/amirhnajafiz/go-graphql/internal/logger"
	"go.uber.org/zap"
)

func Execute() {
	c := config.Load()
	l := logger.New(logger.Config{})
	s, err := gql.Init(c.GQL)

	if err != nil {
		l.Error("schema creation failed", zap.Error(err))
	}

	app := server.Server{
		L: l,
		S: s,
	}.Init()
	_ = app.SetTrustedProxies([]string{c.Proxy})

	err = app.Run(c.Port)
	if err != nil {
		l.Error("server start failed", zap.Error(err))
	}
}
