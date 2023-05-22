// router.go
package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func setupRoutes(router *gin.Engine, db *gorm.DB) {
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

		if err := db.Where("username = ?", username).First(&user).Error; err != nil {
			c.JSON(200, gin.H{"error": "用户名错误"})
			return
		}

		// 比对密码
		if user.Password != password {
			c.JSON(200, gin.H{"error": "密码错误"})
			return
		}

		// 登录成功，返回成功的JSON响应
		tokenString, err := generateToken(username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": tokenString})
	})

	auth := router.Group("/auth")
	auth.Use(validateToken)
	{
		// 添加需要鉴权的路由
		auth.GET("/protected", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Access granted to protected route"})
		})
	}
}
