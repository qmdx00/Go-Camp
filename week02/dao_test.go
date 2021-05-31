package week02

import (
	"database/sql"
	"github.com/pkg/errors"
	"log"
	"testing"
)

func TestHello(t *testing.T) {
	dao := &XXXDao{}
	res, err := dao.GetRow()
	if errors.Is(err, sql.ErrNoRows) {
		log.Fatal("record not found")
	}
	log.Println(res)
}
