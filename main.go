// main.go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// 连接到数据库
	db, err := connectDatabase()
	if err != nil {
		panic("无法连接到数据库")
	}

	// 创建一个默认的gin路由引擎
	router := gin.Default()

	// 设置路由
	setupRoutes(router, db)

	// 设置静态文件目录
	router.Static("/static", "./static")

	// 运行应用程序
	if err := router.Run(":8080"); err != nil {
		panic("无法启动服务器")
	}
}
