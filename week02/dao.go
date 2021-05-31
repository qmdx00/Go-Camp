package week02

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"log"
)

type XXXDao struct{}

type Result struct{}

func (*XXXDao) GetRow() (*Result, error) {
	db, err := sql.Open("mysql", "root:123@tcp(127.0.0.1:3306)/hello")
	if err != nil {
		log.Fatal("open mysql error")
	}

	res := &Result{}
	s := "select * from test"

	err = db.QueryRow(s).Scan(&res)
	if err == sql.ErrNoRows {
		return nil, errors.Wrap(err, fmt.Sprintf("[SQL]: %s,  [Error]: %v", s, err))
	}
	return res, nil
}
