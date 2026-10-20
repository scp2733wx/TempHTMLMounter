package service

import (
	"temphtmlmounter/dao"
	"temphtmlmounter/model"
	"temphtmlmounter/pkg/errcode"
	"temphtmlmounter/pkg/hash"
	"temphtmlmounter/pkg/jwtutil"
)

// Register 注册新用户
func Register(req *model.RegisterRequest) (*model.User, error) {
	exist, err := dao.GetUserByUsername(req.Username)
	if err != nil {
		return nil, errcode.ErrInternalServer
	}
	if exist != nil {
		return nil, errcode.ErrUserExist
	}
	passHash, err := hash.Hash(req.Password)
	if err != nil {
		return nil, errcode.ErrInternalServer
	}
	user := &model.User{
		UserName: req.Username,
		PassHash: passHash,
		NickName: req.Nickname,
		Phone:    req.Phone,
		Email:    req.Email,
		Role:     model.RoleStudent,
	}
	if err := dao.CreateUser(user); err != nil {
		return nil, errcode.ErrInternalServer
	}
	return user, nil
}

// Login 校验用户名密码并签发 token
func Login(req *model.LoginRequest) (*model.LoginResponse, error) {
	user, err := dao.GetUserByUsername(req.Username)
	if err != nil {
		return nil, errcode.ErrInternalServer
	}
	if user == nil || hash.CheckHash(user.PassHash, req.Password) != nil {
		return nil, errcode.ErrUserPwdWrong
	}
	accessToken, err := jwtutil.GenerateAccessToken(user.ID, user.Role)
	if err != nil {
		return nil, errcode.ErrInternalServer
	}
	refreshToken, err := jwtutil.GenerateRefreshToken(user.ID, user.Role)
	if err != nil {
		return nil, errcode.ErrInternalServer
	}
	return &model.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    2 * 3600,
		UserInfo:     *toUserInfo(user),
	}, nil
}

// GetUserInfo 获取用户信息
func GetUserInfo(userID uint) (*model.UserInfo, error) {
	user, err := dao.GetUserByID(userID)
	if err != nil {
		return nil, errcode.ErrInternalServer
	}
	if user == nil {
		return nil, errcode.ErrUserNotExist
	}
	return toUserInfo(user), nil
}

func toUserInfo(u *model.User) *model.UserInfo {
	return &model.UserInfo{
		UserID:     u.ID,
		UserName:   u.UserName,
		NickName:   u.NickName,
		Avatar:     u.Avatar,
		Phone:      u.Phone,
		Email:      u.Email,
		Role:       u.Role,
		CreateTime: u.CreatedAt,
	}
}