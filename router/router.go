package router

import (
	"temphtmlmounter/handler"
	"temphtmlmounter/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	api := r.Group("/api/v1")

	// 健康检查
	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"code": 200, "msg": "pong"})
	})

	// 认证接口（公开）
	auth := api.Group("/auth", middleware.AuthLimit)
	{
		auth.POST("/register", handler.Register)
		auth.POST("/login", handler.Login)
	}

	// 需登录接口
	protected := api.Group("", middleware.JWTAuthMiddleware(), middleware.RequireAuthMiddleware())
	{
		authMe := protected.Group("/auth")
		authMe.GET("/me", handler.GetMe)
	}
}