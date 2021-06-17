package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"week04/app/book/internal/biz"
	"week04/app/book/internal/data/ent"
	"week04/app/book/internal/pkg"
)

type BookService struct {
	Client *ent.Client
	Biz    *biz.BookBusiness
}

// Todo grpc接口测试，makefile编写，dockerfile 和 docker-compose

func (*BookService) CreateBook(ctx *gin.Context) {
	// Todo
}

func (b *BookService) GetBookById(ctx *gin.Context) {
	id := pkg.StringToInt(ctx.Param("id"))
	book, err := b.Biz.QueryBookById(context.Background(), b.Client, id)
	// Todo 异常封装和处理，错误码设计
	if err != nil {
		ctx.JSON(http.StatusNotFound, err)
	}
	ctx.JSON(http.StatusOK, book)
}

func (*BookService) UpdateBookById(ctx *gin.Context) {
	// Todo
}

func (*BookService) DeleteBookById(ctx *gin.Context) {
	// Todo
}
