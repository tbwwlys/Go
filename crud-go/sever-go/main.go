package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func main() {
	fmt.Println("hello world")
	// mysql 地址
	dsn := "root:123456@tcp(127.0.0.1:3306)/crud?charset=utf8mb4&parseTime=True&loc=Local"
	// 链接了数据库 连接了 模型的对应关系
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 解决表名复数问题
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	sqlDB, err := db.DB()
	// 最大空闲数
	sqlDB.SetMaxIdleConns(10)
	// 最大链接数
	sqlDB.SetMaxOpenConns(100)
	// 10 秒超时
	sqlDB.SetConnMaxLifetime(10000)

	type List struct {
		gorm.Model
		ID      uint8  `json:"id"`
		Name    string `json:"name"`
		Age     uint8  `json:"age"`
		School  string `json:"school"`
		Phone   string `json:"phone"`
		Address string `json:"address"`
	}

	db.AutoMigrate(&List{})
	fmt.Println("db:", db)
	fmt.Println("err:", err)
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		// default /
		c.JSON(200, gin.H{
			"msg":  "hello world",
			"code": 0,
		})
	})
	PORT := "8081"
	r.Run(":" + PORT)
}
