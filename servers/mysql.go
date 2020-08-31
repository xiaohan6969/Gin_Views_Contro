package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var SqlDB *sql.DB

func init() {
	var err error
	//practice1 是数据库的名字，用户名：root:mysql是密码
	SqlDB, err = sql.Open("mysql", "root:123456@tcp(122.51.54.111:3306)/practice1")
	if err != nil {
		fmt.Println("err1==",err)
	}
	err = SqlDB.Ping()
	if err != nil {
		fmt.Println("err2==",err)
	}
}
