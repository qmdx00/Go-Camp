package data

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"week04/app/book/internal/conf"
	"week04/app/book/internal/data/ent"
)

// ProvideSet for data package ...
var ProvideSet = wire.NewSet(NewEntClient)

func NewEntClient(options conf.Options) *ent.Client {
	client, err := ent.Open(options.Data.Database.Driver, options.Data.Database.DataSource)
	if err != nil {
		panic(err)
	}
	return client
}
