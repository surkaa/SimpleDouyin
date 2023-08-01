package service

import (
	"DouyinBackend/request"
	"DouyinBackend/response"
)

// UserService 用户业务操作层
type UserService interface {

	// Register 注册
	Register(body request.UserRegisterBody) (response.UserRegister, error)

	// Login 登录
	Login(body request.UserLoginBody) (response.UserLogin, error)

	// Info 获取用户信息
	Info(body request.UserInfoBody) (response.UserInfo, error)
}
