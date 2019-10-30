package main

import (
	router "Gin_Views_Contro/routers"
	db "Gin_Views_Contro/servers"
	"fmt"
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
	fmt.Println(r.Run(":1234").Error())

}
