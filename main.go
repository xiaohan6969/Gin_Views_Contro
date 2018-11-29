package main

import (
	db "Bokeyuan/database"
	router "Bokeyuan/routers"
)

func main() {
	//数据库
	defer db.SqlDB.Close()

	//路由部分
	router := router.InitRouter()

	//静态资源
	router.Static("/static", "./static")

	//运行的端口
	router.Run(":9869")

}
