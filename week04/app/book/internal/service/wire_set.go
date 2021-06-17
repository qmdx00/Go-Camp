package service

import (
	"github.com/google/wire"
	"week04/app/book/internal/biz"
	"week04/app/book/internal/data/ent"
)

// ProvideSet for service package ...
var ProvideSet = wire.NewSet(NewBookService)

func NewBookService(client *ent.Client, business *biz.BookBusiness) *BookService {
	return &BookService{Client: client, Biz: business}
}
