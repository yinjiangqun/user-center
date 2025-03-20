package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/yinjiangqun/user-center/backend/config"
	"github.com/yinjiangqun/user-center/backend/routers"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const BACKEND_SERVER_PORT = 8081
var DB *gorm.DB

func InitDB() {
	dsn := "root:King@01010101@tcp(127.0.0.1:3306)/user_center?charset=utf8mb4&parseTime=True&loc=Asia%2FShanghai"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
	}
	
	// 获取底层的sqlDB
	sqlDB, err := DB.DB()
	if err != nil {
			log.Fatalf("Failed to get database: %v", err)
	}
	
	// 配置连接池
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(0)
	
	// 测试连接
	if err := sqlDB.Ping(); err != nil {
			log.Fatalf("Database connection test failed: %v", err)
	}
}
	
func main() {
	// 初始化数据库
	InitDB()

	// 读取配置文件
	cfg := config.GetConfig()  // Get configuration instance
	port := cfg.GetInt("server.port") // Get port number directly from configuration
	fmt.Printf("Server is running on port: %d\n", port)

	router := gin.Default() // 创建 Gin 引擎实例

	// 加载全局中间件
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// 注册路由
	routers.RegisterUserRoutes(router)
	routers.RegisterAdminRoutes(router)

	// 定义简单的 GET 路由
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// 启动服务器，监听 8081 端口
	router.Run(":8081")
}
