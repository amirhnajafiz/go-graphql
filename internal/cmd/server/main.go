package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/amirhnajafiz/go-graphql/internal/gql"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"go.uber.org/zap"
)

type Server struct {
	L *zap.Logger
	S graphql.Schema
}

func (s Server) Init() *gin.Engine {
	app := gin.Default()

	gin.SetMode(gin.TestMode)

	app.GET("/", func(context *gin.Context) {
		s.L.Info("root request")

		context.HTML(http.StatusOK, "template/index.html", gin.H{})
	})

	app.POST("/query", func(context *gin.Context) {
		jsonData, err := ioutil.ReadAll(context.Request.Body)
		if err != nil {
			_ = context.Error(err)

			s.L.Error("query read failed", zap.Error(err))

			return
		}

		var query struct {
			Query string `json:"query"`
		}

		err = json.Unmarshal(jsonData, &query)
		if err != nil {
			_ = context.Error(err)

			s.L.Error("unmarshalling failed", zap.Error(err))

			return
		}

		r := gql.ExecuteQuery(query.Query, s.S)
		rJSON, _ := json.Marshal(r)

		s.L.Info("successful query executed")

		context.JSON(http.StatusOK, rJSON)
	})

	return app
}
