package main

import (
	router "Gin_Views_Contro/routers"
	db "Gin_Views_Contro/servers"
)

func main() {
	//数据库
	defer db.SqlDB.Close()
	//gin.SetMode("test")
	//路由部分
	r:= router.InitRouter()

	//静态资源
	r.Static("/static", "./static")

	//运行的端口
	panic(r.Run(":1234").Error())

}
