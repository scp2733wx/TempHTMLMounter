package handler

import (
	"temphtmlmounter/model"
	"temphtmlmounter/pkg/errcode"
	"temphtmlmounter/pkg/response"
	"temphtmlmounter/service"

	"github.com/gin-gonic/gin"
)

// Register 注册
func Register(c *gin.Context) {
	var req model.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Err(c, errcode.ErrBadRequest)
		return
	}
	user, err := service.Register(&req)
	if err != nil {
		response.Err(c, err)
		return
	}
	response.OK(c, user)
}

// Login 登录
func Login(c *gin.Context) {
	var req model.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Err(c, errcode.ErrBadRequest)
		return
	}
	data, err := service.Login(&req)
	if err != nil {
		response.Err(c, err)
		return
	}
	response.OK(c, data)
}

// GetMe 获取当前登录用户信息
func GetMe(c *gin.Context) {
	userID := c.GetUint("id")
	if userID == 0 {
		response.Err(c, errcode.ErrUnauthorized)
		return
	}
	info, err := service.GetUserInfo(userID)
	if err != nil {
		response.Err(c, err)
		return
	}
	response.OK(c, info)
}