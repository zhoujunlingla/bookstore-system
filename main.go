package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 创建一个模型结构体，用于表示用户信息
type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

func main() {
	// 连接到数据库
	dsn := "root:Zjl7758258@tcp(localhost:3306)/bookstore?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("无法连接到数据库")
	}
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	// 迁移模型结构到数据库
	db.AutoMigrate(&User{})

	// 创建一个默认的gin路由引擎
	router := gin.Default()

	// 设置登录页面的路由
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	// 处理登录请求的路由
	router.POST("/login", func(c *gin.Context) {
		// 从请求中获取用户名和密码
		username := c.PostForm("username")
		password := c.PostForm("password")

		// 在数据库中查找用户
		var user User
		if err := db.Where("username = ? AND password = ?", username, password).First(&user).Error; err != nil {
			c.JSON(200, gin.H{"error": "用户名或密码错误"})
			return
		}

		// 登录成功，返回成功的JSON响应
		c.JSON(200, gin.H{"message": "登录成功"})
	})

	// 设置静态文件目录
	router.Static("/static", "./static")

	// 运行应用程序
	if err := router.Run(":8080"); err != nil {
		panic("无法启动服务器")
	}
}
