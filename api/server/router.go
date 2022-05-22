package server

import (
	"github.com/georgejdanforth/crate-digger/server/handlers"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	{
		api.GET("/healthz", func (ctx *gin.Context) {
			ctx.Data(200, gin.MIMEJSON, nil)
		})

		v1 := api.Group("/v1")
		{
			v1.GET("/search", handlers.Search)
		}
	}

	return r
}
