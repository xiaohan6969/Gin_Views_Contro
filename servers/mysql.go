package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var SqlDB *sql.DB

func init() {
	var err error
	//practice1 是数据库的名字，用户名：root:mysql是密码
	SqlDB, err = sql.Open("mysql", "root:pwd12345@tcp(122.51.54.111:3306)/practice1")
	if err != nil {
		log.Fatal(err.Error())
	}
	err = SqlDB.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
}
