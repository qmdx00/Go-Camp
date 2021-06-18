package service

import (
	"github.com/google/wire"
	"week04/app/book/internal/biz"
	"week04/app/book/internal/data/ent"
)

// ProvideSet for service package ...
var ProvideSet = wire.NewSet(NewHTTPBookService)

func NewHTTPBookService(client *ent.Client) biz.HTTPBookRepo {
	return &HTTPBookService{Client: client}
}

//func NewGRPCBookService(client *ent.Client) biz.GRPCBookRepo {
//	return &GRPCBookService{Client: client}
//}
