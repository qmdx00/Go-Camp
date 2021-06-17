package biz

import "github.com/google/wire"

// ProvideSet for biz package ...
var ProvideSet = wire.NewSet(NewBookBusiness)

func NewBookBusiness() *BookBusiness {
	return &BookBusiness{}
}
