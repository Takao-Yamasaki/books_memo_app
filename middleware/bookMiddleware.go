package service

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func Env_load() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("could not read .env file")
	}
}

func DbConnect() *gorm.DB {
	Env_load()

	DBMS := "mysql"
	USER := os.Getenv("MYSQL_USER")
	PASS := os.Getenv("MYSQL_PASSWORD")
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "book"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}
