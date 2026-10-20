package errcode

type BizError struct {
	HttpStatus int
	Code       int
	Message    string
}

func (e *BizError) Error() string {
	return e.Message
}

func New(httpStatus int, code int, msg string) *BizError {
	return &BizError{
		HttpStatus: httpStatus,
		Code:       code,
		Message:    msg,
	}
}

var (
	ErrSuccess          = &BizError{200, 200, "success"}
	ErrBadRequest       = &BizError{400, 1001, "参数错误"}
	ErrUnauthorized     = &BizError{401, 1002, "未登录"}
	ErrForbidden        = &BizError{403, 1003, "权限不足"}
	ErrResourceNotFound = &BizError{404, 1004, "资源不存在"}
	ErrTooManyRequests  = &BizError{429, 1005, "请求过于频繁"}
	ErrTokenExpired     = &BizError{401, 1006, "登录已过期"}

	ErrUserExist       = &BizError{409, 2001, "用户名已存在"}
	ErrUserPwdWrong    = &BizError{400, 2002, "用户名或密码错误"}
	ErrOldPwdWrong     = &BizError{400, 2004, "原密码错误"}
	ErrUserNotExist    = &BizError{404, 2008, "用户不存在"}

	ErrInternalServer = &BizError{500, 5000, "服务器内部错误"}
)