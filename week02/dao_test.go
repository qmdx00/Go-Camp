package week02

import (
	"database/sql"
	"github.com/pkg/errors"
	"testing"
)

func TestHello(t *testing.T) {
	dao := &XXXDao{}
	res, err := dao.GetRow()
	if errors.Is(err, sql.ErrNoRows) {
		t.Fatal("record not found")
	}
	t.Log(res)
}
