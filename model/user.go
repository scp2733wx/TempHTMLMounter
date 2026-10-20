package model

import (
	"time"

	"gorm.io/gorm"
)

// 用户角色常量
const (
	RoleStudent = "student"
	RoleAdmin   = "admin"
)

type User struct {
	ID        uint           `json:"userId"`
	UserName  string         `gorm:"size:20;uniqueIndex" json:"username"`
	PassHash  string         `gorm:"size:255" json:"-"`
	NickName  string         `gorm:"size:20" json:"nickname"`
	Phone     string         `gorm:"size:11" json:"phone"`
	Email     string         `gorm:"size:64" json:"email"`
	Avatar    string         `gorm:"size:255" json:"avatar"`
	Role      string         `gorm:"size:16;default:student" json:"role"` // student | admin
	CreatedAt time.Time      `json:"createTime"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=4,max=20"`
	Password string `json:"password" binding:"required,min=8,max=20"`
	Nickname string `json:"nickname" binding:"required,min=1,max=20"`
	Phone    string `json:"phone" binding:"required,len=11"`
	Email    string `json:"email" binding:"required,email"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToken  string   `json:"accessToken"`
	RefreshToken string   `json:"refreshToken"`
	ExpiresIn    int64    `json:"expiresIn"`
	UserInfo     UserInfo `json:"userInfo"`
}

type UserInfo struct {
	UserID     uint      `json:"userId"`
	UserName   string    `json:"username"`
	NickName   string    `json:"nickname"`
	Avatar     string    `json:"avatar"`
	Phone      string    `json:"phone"`
	Email      string    `json:"email"`
	Role       string    `json:"role"`
	CreateTime time.Time `json:"createTime"`
}