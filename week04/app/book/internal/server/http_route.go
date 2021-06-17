package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"week04/app/book/internal/conf"
)

// NewHttpRouter set gin.Engine as http.Handler
func NewHttpRouter(options conf.Options) *gin.Engine {
	gin.SetMode(options.Mode)
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello world")
	})

	return router
}
