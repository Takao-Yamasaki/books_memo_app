package service

import (
	"books_memo_app/model"
	"database/sql/driver"
	"errors"
	"fmt"
	"log"

	"github.com/go-xorm/xorm"
)

var DbEngine *xorm.Engine

func init() {
	driverName := "mysql"
	DsName := "root:password@(localhost:3306)/book?charset=utf8"
	err := errors.New("")
	DbEngine, err = xorm.NewEngine(driverName, DsName)
	if err != nill && err.Error() != "" {
		panic(err.Error())
	}
	DbEngine.ShowSQL(true)
	DbEngine.SetMaxOpenConns(2)
	DbEngine.Sysc2(new(model.book))
	fmt.Println("init data base ok")
}
