package impl

import (
	"DouyinBackend/request"
	"DouyinBackend/response"
)

// UserServiceImpl 用户业务层实现类
type UserServiceImpl struct {
}

// Register 注册
func (userServiceImpl UserServiceImpl) Register(body request.UserRegisterBody) (response.UserRegister, error) {
	return response.UserRegister{}, nil
}

// Login 登录
func (userServiceImpl UserServiceImpl) Login(body request.UserLoginBody) (response.UserLogin, error) {
	return response.UserLogin{}, nil
}

// Info 获取用户信息
func (userServiceImpl UserServiceImpl) Info(body request.UserInfoBody) (response.UserInfo, error) {
	return response.UserInfo{}, nil
}
