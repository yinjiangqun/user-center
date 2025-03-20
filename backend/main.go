package main

import (

	//"github.com/yinjiangqun/user-center/cmd"
	"log"
	//"github.com/yinjiangqun/user-center/backend/config.json"
	"backend/config.json"

	"github.com/gin-gonic/gin"

	//"github.com/yinjiangqun/user-center/backend/routers"
	"backend/routers"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	//"github.com/yinjiangqun/user-center/backend/config"
	"backend/config"
)

const BACKEND_SERVER_PORT = 8081
var DB *gorm.DB

func InitDB() {
	dsn := "host=127.0.0.1 user=postgres password=king01010101 dbname=user_center port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
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

	// 读取配置文件
	config := config.Config()
	port := config.Config().GetInt("server.port") // 获取端口号
	fmt.Printf("Server is running on port: %d\n", port)
	// 启动服务器，监听 8080 端口
	router.Run(":8080")
}
