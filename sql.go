// sql.go
package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 创建一个模型结构体，用于表示用户信息
type User struct {
	gorm.Model
	Username string `gorm:"username"`
	Password string `gorm:"password"`
}

func connectDatabase() (*gorm.DB, error) {
	dsn := "root:Zjl7758258@tcp(localhost:3306)/bookstore?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	err = sqlDB.Ping()
	if err != nil {
		return nil, err
	}

	// 迁移模型结构到数据库
	err = db.AutoMigrate(&User{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
