package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/amirhnajafiz/go-graphql/internal/gql"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

func Init(s graphql.Schema) *gin.Engine {
	app := gin.Default()

	app.POST("/query", func(context *gin.Context) {
		jsonData, err := ioutil.ReadAll(context.Request.Body)
		if err != nil {
			_ = context.Error(err)
			return
		}

		var query struct {
			Query string `json:"query"`
		}

		err = json.Unmarshal(jsonData, &query)
		if err != nil {
			_ = context.Error(err)
			return
		}

		r := gql.ExecuteQuery(query.Query, s)
		rJSON, _ := json.Marshal(r)

		context.JSON(http.StatusOK, rJSON)
	})

	return app
}
