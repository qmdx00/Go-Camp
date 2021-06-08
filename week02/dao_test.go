package week02

import (
	"database/sql"
	"github.com/pkg/errors"
	"testing"
)

func TestHello(t *testing.T) {
	dao := &XXXDao{}
	res, err := dao.GetRow()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			t.Fatal("record not found")
		} else {
			t.Fatal("other error", err)
		}
	}
	t.Log(res)
}
