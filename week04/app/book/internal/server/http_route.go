package server

import (
	"github.com/gin-gonic/gin"
	"week04/app/book/internal/conf"
	"week04/app/book/internal/service"
)

// NewHttpRouter set gin.Engine as http.Handler
func NewHttpRouter(options conf.Options, service *service.BookService) *gin.Engine {
	gin.SetMode(options.Mode)
	router := gin.Default()

	// set book router
	bookRouter(router, service)

	return router
}

func bookRouter(app *gin.Engine, service *service.BookService) {
	v := app.Group("v1/books")

	v.POST("/", service.CreateBook)
	v.GET("/:id", service.GetBookById)
	v.PUT("/:id", service.UpdateBookById)
	v.DELETE("/:id", service.DeleteBookById)
}
