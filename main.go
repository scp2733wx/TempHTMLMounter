package main

import (
	"log"
	"temphtmlmounter/config"
	"temphtmlmounter/dao"
	"temphtmlmounter/middleware"
	"temphtmlmounter/model"
	"temphtmlmounter/pkg/jwtutil"
	"temphtmlmounter/pkg/logger"
	"temphtmlmounter/pkg/redisdb"
	"temphtmlmounter/router"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	dsn := config.GetConfig().Mysql.Dsn
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("获取数据库连接池失败: %v", err)
	}
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)
	return db
}

func main() {
	logger.InitLogger()
	db := initDB()

	// 自动迁移数据库表
	if err := db.AutoMigrate(&model.User{}); err != nil {
		log.Fatalf("数据库迁移失败: %v", err)
	}
	dao.InitDB(db)

	redisdb.InitRedis()

	if err := jwtutil.Init(config.GetConfig().JWT.Secret); err != nil {
		log.Fatalf("JWT初始化失败: %v", err)
	}

	r := gin.Default()
	r.Use(middleware.AccessLog())
	r.Use(middleware.ErrorMiddleware())
	router.InitRouter(r)

	port := config.GetConfig().Server.Port
	log.Printf("服务启动，监听端口: %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}
}